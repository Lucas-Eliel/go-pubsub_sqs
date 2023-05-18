package main

import (
	"log"
	"pubsub_sqs/internal/consumers"
	"pubsub_sqs/internal/models"
	"pubsub_sqs/internal/producers"
	"pubsub_sqs/pkg/sqs"
	"time"
)

func main() {

	clientsqs, queueUrl := sqs.CreateLocalClient()

	user := models.User{
		ID:   "253",
		Nome: "Pedro",
	}

	producers.Producer(clientsqs, queueUrl, user)
	log.Printf("Enviado mensagem com sucesso. Os dados são id: %s e nome: %s", user.ID, user.Nome)

	time.Sleep(10)

	resposta := consumers.Consumer(clientsqs, queueUrl)

	log.Printf("Recebido mensagem com sucesso. Os dados são id: %s e nome: %s", resposta.ID, resposta.Nome)
}
