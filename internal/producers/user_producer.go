package producers

import (
	"context"
	"encoding/json"
	"pubsub_sqs/internal/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Producer(client *sqs.Client, queueUrl *string, user models.User) {

	b, e := json.Marshal(user)

	if e != nil {
		panic(e)
	}

	_, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		MessageBody: aws.String(string(b)),
		QueueUrl:    queueUrl,
	})

	if err != nil {
		panic(err)
	}
}
