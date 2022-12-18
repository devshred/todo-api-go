#!/bin/bash

CONTAINER_NAME=todo-api-go

wait_for_log() {
  _container_name=$1
  _needle=$2
  _seconds=0
  _timeout=60

  while true
  do
    logs=$(docker logs $_container_name 2>&1 | grep "$_needle")
    if [ -n "$logs" ]
    then
      echo $logs
      break
    fi

    if [ $_seconds -ge $_timeout ]
    then
      echo "timeout ${_timeout}s exceeded"
      break
    fi
    sleep 1
  done
}

docker run --rm -d \
    -e DB_HOST=db \
    -e DB_USER=todo \
    -e DB_PASS=todo \
    -e DB_NAME=todo \
    -e DB_PORT=5432 \
    --network=todo \
    -p "8080:8080" \
    --name $CONTAINER_NAME \
    -m 512M \
    todo-api-go

wait_for_log $CONTAINER_NAME "Started Application"

mem_usage=$(docker stats $CONTAINER_NAME --no-stream --format "{{ json .MemUsage }}" | jq . -r | awk '{print $1}')
echo "Memory usage: $mem_usage"
