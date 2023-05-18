#!/bin/bash

echo Encerrando ambiente local

docker compose -f ./sqs/docker-compose.yml down

docker network rm go_pubsub_sqs