package mq

import (
	"encoding/json"

	"github.com/Levap123/playstar-test/logging_service/entities"
	"github.com/Levap123/playstar-test/logging_service/logs"
	"github.com/streadway/amqp"
)

type Consumer struct {
	delivery <-chan amqp.Delivery
	logger   *logs.Logger
	repo     ILogsRepo
}

type ILogsRepo interface {
	Insert(logMsg entities.LogMessage)
}

func NewConsumer(mqChan *amqp.Channel, logger *logs.Logger, repo ILogsRepo) (*Consumer, error) {
	messages, err := mqChan.Consume(
		"logs",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		delivery: messages,
		logger:   logger,
		repo:     repo,
	}, nil
}

func (c *Consumer) Consume() {
	c.logger.Info().Msg("started consuming messages from queue...")

	var logMsg entities.LogMessage
	for msg := range c.delivery {

		if err := json.Unmarshal(msg.Body, &logMsg); err != nil {
			c.logger.Err(err).Msg("error in unmarshalling message from queue")
			continue
		}

		go c.repo.Insert(logMsg)
	}
}
