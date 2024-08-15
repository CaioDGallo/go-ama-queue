for i in {1..10}; do sleep $(shuf -i 1-5 -n 1);./stress_test.sh; done
