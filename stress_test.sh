DOMAIN=$1

#check if domain is null or empty string
if [ -z "$DOMAIN" ]; then
    DOMAIN="http://localhost:8080"
fi

for i in $(seq 10 $CYCLES); do
    DURATION="-d$(shuf -i 1-10 -n 1)s"
    CONNECTIONS=$(shuf -i 10-50 -n 1)
    THREADS=$(shuf -i 1-12 -n 1)
    CYCLES=$(shuf -i 11-100 -n 1)
    SLEEP_TIME=$(shuf -i 1-10 -n 1)

    echo "Running test $i"
    wrk -t$THREADS -c$CONNECTIONS $DURATION -s stress_test_create_room_body.lua "$DOMAIN/api/rooms"
    sleep $SLEEP_TIME
done
