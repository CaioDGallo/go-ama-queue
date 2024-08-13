wrk -t12 -c400 -d10s -s queue_stress_test.lua http://localhost:8080/api/rooms
