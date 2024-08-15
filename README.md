# Go - GAMA Message Queue and Metrics

![GoLang + RabbitMQ + Prometheus + Grafana](docs/assets/hero.png)

![License](https://img.shields.io/github/license/CaioDGallo/go-ama)

Welcome to the GAMA (Go - Ask Me Anything) async job queue. This project is built using GoLang, RabbitMQ, Prometheus, and Grafana.
## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Introduction

The GAMA Queue asynchronously processes user data collected from the main GAMA application. RabbitMQ is used as the message broker, Prometheus + Grafana for metrics observability and dashboards, and GoLang leveraging the power of goroutines to concurrently process the messages.


## Features

- **Asynchronous Job Queue:** The GAMA Queue processes user data asynchronously using RabbitMQ.
- **Metrics and Observability:** Prometheus and Grafana are used to monitor the application's performance and health.
- **Concurrent Processing:** GoLang's goroutines are used to concurrently process the messages.
- **Dockerized Environment:** The application is dockerized for easy setup and deployment.

## Installation

To get started with the Go AMA application, follow these steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/CaioDGallo/go-ama-queue.git
   cd go-ama-queue
   ```

2. **Build the Docker environment:**

   ```bash
   docker compose up
   ```

3. **The application should be up and running with all its dependencies:**

   ```
   STDIN: Queue and general logging
   METRICS: http://localhost:8081/metrics
   ```

## Usage

Once the application is up and running, you can start processing messages. Currently the only two endpoints that generate user data to be processed are the Create Room and the Ask a Question endpoints:

This is the postman collection to run the endpoints of the main [GAMA](https://github.com/CaioDGallo/go-ama) application:
[![Run in Postman](https://run.pstmn.io/button.svg)](https://github.com/CaioDGallo/go-ama/docs/GAMA.postman_collection.json)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

Thank you for checking out the GAMA Queue application! If you have any questions or need further assistance, feel free to reach out.

