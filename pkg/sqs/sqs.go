package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func CreateLocalClient() (*sqs.Client, *string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),

		config.WithRegion("sa-east-1"),

		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:4566", SigningRegion: "sa-east-1"}, nil
			})),

		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "admin", SecretAccessKey: "admin",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	client := sqs.NewFromConfig(cfg)

	gQInput := &sqs.GetQueueUrlInput{
		QueueName: aws.String("user-message"),
	}

	result, erro := client.GetQueueUrl(context.TODO(), gQInput)

	if erro != nil {
		panic(erro)
	}

	return client, result.QueueUrl

}
