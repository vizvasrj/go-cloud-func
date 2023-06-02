#!/bin/bash

TOTAL_REQUESTS=10
CONCURRENT_REQUESTS=5
URL="https://go-cloud-func-5wxqelilkq-uc.a.run.app/ping"

for ((i=1; i<=$TOTAL_REQUESTS; i++)); do
    hey -n $CONCURRENT_REQUESTS -c $CONCURRENT_REQUESTS $URL
done