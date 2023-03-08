package mq

import "github.com/streadway/amqp"

func InitConnection(addr string) (*amqp.Connection, error) {
	connRabbitMQ, err := amqp.Dial(addr)
	if err != nil {
		return nil, err
	}

	return connRabbitMQ, nil
}
