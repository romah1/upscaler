package common

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type ConnectionHolder struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func NewConnectionHolder(url string, params QueueParams) (*ConnectionHolder, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		params.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	return &ConnectionHolder{
		Connection: conn,
		Channel:    ch,
		Queue:      q,
	}, nil
}

func (sender *ConnectionHolder) Close() (channelErr error, connectionErr error) {
	channelErr = sender.Channel.Close()
	connectionErr = sender.Connection.Close()
	return
}
