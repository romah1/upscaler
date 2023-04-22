package receiver

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"upscaler/tasks_queue/common"
)

type Receiver struct {
	connectionHolder *common.ConnectionHolder
	queue            *amqp.Queue
}

func NewSender(url string, params common.QueueParams) (*Receiver, error) {
	mqConnectionHolder, err := common.NewConnectionHolder(url, params)
	if err != nil {
		return nil, err
	}

	return &Receiver{
		connectionHolder: mqConnectionHolder,
		queue:            nil,
	}, nil
}

func (receiver *Receiver) Close() (channelErr error, connectionErr error) {
	channelErr, connectionErr = receiver.connectionHolder.Close()
	return
}

func (receiver *Receiver) Receive() (<-chan amqp.Delivery, error) {
	delivery, err := receiver.connectionHolder.Channel.Consume(
		receiver.connectionHolder.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return delivery, err
}
