#!/bin/bash

docker build -t todo-api-go:latest .
docker tag todo-api-go:latest registry.local:5000/todo-api-go:latest
docker push registry.local:5000/todo-api-go:latest