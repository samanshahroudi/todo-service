package sqs

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/samanshahroudi/todo-service/internal/domain"
	"github.com/samanshahroudi/todo-service/internal/ports"
)

type SQSAdapter struct {
	Client   *sqs.Client
	QueueURL string
}

func NewSQSAdapter(client *sqs.Client, queueURL string) ports.SQSService {
	return &SQSAdapter{
		Client:   client,
		QueueURL: queueURL,
	}
}

func (s *SQSAdapter) SendMessage(todo domain.TodoItem) error {
	body, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	_, err = s.Client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    aws.String(s.QueueURL),
		MessageBody: aws.String(string(body)),
	})

	return err
}
