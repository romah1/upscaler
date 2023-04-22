package publisher

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"upscaler/tasks_queue/common"
)

type Publisher struct {
	connectionHolder *common.ConnectionHolder
	queue            *amqp.Queue
}

func NewSender(url string, params common.QueueParams) (*Publisher, error) {
	mqConnectionHolder, err := common.NewConnectionHolder(url, params)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		connectionHolder: mqConnectionHolder,
		queue:            nil,
	}, nil
}

func (publisher *Publisher) Close() (channelErr error, connectionErr error) {
	channelErr, connectionErr = publisher.connectionHolder.Close()
	return
}

func (publisher *Publisher) Publish(publishing amqp.Publishing) error {
	err := publisher.connectionHolder.Channel.PublishWithContext(
		context.TODO(),
		"",
		publisher.connectionHolder.Queue.Name,
		false,
		false,
		publishing,
	)
	if err != nil {
		return err
	}
	return nil
}
