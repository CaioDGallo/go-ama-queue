package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"

	"github.com/CaioDGallo/go-ama-queue/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

const (
	QueueCollectUserDataName = "collect_user_data"
)

const (
	CollectedUserDataActionCreateRoom    = "create_room"
	CollectedUserDataActionCreateMessage = "create_message"
)

type UserDataMessage struct {
	IP               string `json:"ip"`
	Location         string `json:"location"`
	Device           string `json:"device"`
	UserAgent        string `json:"user_agent"`
	Action           string `json:"action"`
	JSONResponseBody string `json:"response_body"`
	Referrer         string `json:"referrer"`
	RequestMethod    string `json:"request_method"`
	RequestPath      string `json:"request_path"`
	Timestamp        string `json:"timestamp"`
}

// Job represents a job to be processed
type Job struct {
	Payload UserDataMessage `json:"payload"`
	ID      uuid.UUID       `json:"id"`
}

// Worker represents a worker that processes jobs
type Worker struct {
	JobChannel chan Job
	WaitGroup  *sync.WaitGroup
	ID         int
}

// NewWorker creates a new worker
func NewWorker(id int, jobChannel chan Job, wg *sync.WaitGroup) Worker {
	return Worker{
		ID:         id,
		JobChannel: jobChannel,
		WaitGroup:  wg,
	}
}

// Start starts the worker to process jobs
func (w Worker) Start(qr *pgstore.Queries) {
	go func() {
		for job := range w.JobChannel {
			startTime := time.Now() // Record start time
			fmt.Printf(
				"Worker %d processing job %s with payload: %s\n",
				w.ID,
				job.ID.String(),
				job.Payload,
			)

			qr.InsertUserData(context.Background(), pgstore.InsertUserDataParams{
				Ip:               job.Payload.IP,
				UserAgent:        job.Payload.UserAgent,
				Location:         job.Payload.Location,
				Device:           job.Payload.Device,
				Action:           job.Payload.Action,
				JsonResponseBody: job.Payload.JSONResponseBody,
				Referrer:         job.Payload.Referrer,
				RequestMethod:    job.Payload.RequestMethod,
				RequestPath:      job.Payload.RequestPath,
			})

			duration := time.Since(startTime) // Calculate duration
			fmt.Printf("Worker %d finished job %s in %v\n", w.ID, job.ID.String(), duration)
			w.WaitGroup.Done()
		}
	}()
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s",
			os.Getenv("GAMA_QUEUE_DATABASE_USER"),
			os.Getenv("GAMA_QUEUE_DATABASE_PASSWORD"),
			os.Getenv("GAMA_QUEUE_DATABASE_HOST"),
			os.Getenv("GAMA_QUEUE_DATABASE_PORT"),
			os.Getenv("GAMA_QUEUE_DATABASE_NAME"),
		))
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	qr := pgstore.New(pool)

	// Set the number of OS threads to use
	runtime.GOMAXPROCS(runtime.NumCPU())

	const numWorkers = 64
	const queueSize = 1000

	jobChannel := make(chan Job, queueSize)
	var wg sync.WaitGroup
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Create and start workers
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i, jobChannel, &wg)
		worker.Start(qr)
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		QueueCollectUserDataName, // name
		true,                     // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for {
			select {
			case d, ok := <-msgs:
				if !ok {
					return
				}
				var job Job
				if err := json.Unmarshal(d.Body, &job); err != nil {
					log.Printf("Failed to decode job: %v", err)
					continue
				}
				wg.Add(1)
				func() {
					defer func() {
						if r := recover(); r != nil {
							log.Printf("Failed to send job to jobChannel: %v", r)
						}
					}()
					jobChannel <- job
				}()
			case <-quit:
				return
			}
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")
	<-quit
	wg.Wait()
	close(jobChannel)

	log.Println("All jobs processed")
}
