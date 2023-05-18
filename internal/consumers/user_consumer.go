package consumers

import (
	"context"
	"encoding/json"
	"pubsub_sqs/internal/models"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Consumer(client *sqs.Client, queueUrl *string) models.User {

	out, err := client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            queueUrl,
		MaxNumberOfMessages: 1,
	})

	if err != nil {
		panic(err)
	}

	message := out.Messages[0]

	response := string(*message.Body)

	user := models.User{}

	json.Unmarshal([]byte(response), &user)

	return user
}
