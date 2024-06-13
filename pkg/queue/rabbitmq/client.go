package rabbitmq

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue"
	"log"
)

type ClientImpl struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
}

var _ queue.Client = (*ClientImpl)(nil)

func NewClient(url string) *ClientImpl {
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	return &ClientImpl{
		connection: conn,
	}
}

func (c ClientImpl) Publish(queue string, body []byte) error {
	err := c.getChannel().Publish(
		amqp091.DefaultExchange,
		queue,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish a message. Error: %s", err)
	}

	log.Printf(" [x] Sent to %s: %s\n", queue, body)

	return err
}

func (c ClientImpl) CreateQueue(name string) amqp091.Queue {
	q, err := c.getChannel().QueueDeclare(name, true, true, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	return q
}

func (c ClientImpl) getChannel() *amqp091.Channel {
	if c.channel != nil && !c.channel.IsClosed() {
		return c.channel
	}

	ch, err := c.connection.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
	}

	c.channel = ch

	return c.channel
}
