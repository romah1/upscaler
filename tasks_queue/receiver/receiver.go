package sender

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"upscaler/tasks_queue/common"
)

type Sender struct {
	connectionHolder *common.ConnectionHolder
	queue            *amqp.Queue
}

func NewSender(url string, params common.QueueParams) (*Sender, error) {
	mqConnectionHolder, err := common.NewConnectionHolder(url, params)
	if err != nil {
		return nil, err
	}

	return &Sender{
		connectionHolder: mqConnectionHolder,
		queue:            nil,
	}, nil
}

func (sender *Sender) Close() (channelErr error, connectionErr error) {
	channelErr, connectionErr = sender.connectionHolder.Close()
	return
}
