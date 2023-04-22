package mq_receiver

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"upscaler/message_queue/mq_common"
)

type Receiver struct {
	connectionHolder *mq_common.ConnectionHolder
	queue            *amqp.Queue
}

func NewReceiver(url string, params mq_common.QueueParams) (*Receiver, error) {
	mqConnectionHolder, err := mq_common.NewConnectionHolder(url, params)
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
