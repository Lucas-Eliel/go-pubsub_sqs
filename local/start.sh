#!/bin/bash

echo Iniciando ambiente local

docker network create go_pubsub_sqs

docker compose -f ./sqs/docker-compose.yml up -d 
