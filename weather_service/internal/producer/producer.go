package producer

import (
	"github.com/Levap123/playstar-test/weather_service/internal/logs"
	"github.com/nxadm/tail"
	"github.com/streadway/amqp"
)

type Producer struct {
	channel *amqp.Channel
	logger  *logs.Logger
}

func NewProducer(channel *amqp.Channel, logger *logs.Logger) (*Producer, error) {

	_, err := channel.QueueDeclare(
		"logs",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &Producer{
		channel: channel,
	}, nil
}

func (p *Producer) Produce() {
	t, err := tail.TailFile("logs.log", tail.Config{Follow: true})
	if err != nil {
		p.logger.Err(err).Msg("unable to read from log file")
	}

	for line := range t.Lines {
		msg := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(line.Text),
		}

		if err := p.channel.Publish(
			"",     // exchange
			"logs", // queue name
			false,  // mandatory
			false,  // immediate
			msg); err != nil {
			p.logger.Err(err).Msg("error in publish message to mq")
		}
	}

}
