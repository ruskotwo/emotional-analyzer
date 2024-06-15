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

func (c ClientImpl) Publish(key string, body []byte) error {
	err := c.getChannel().Publish(
		amqp091.DefaultExchange,
		key,
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

	log.Printf("Sent to %s: %s\n", key, body)

	return err
}

func (c ClientImpl) CreateQueue(key string) amqp091.Queue {
	q, err := c.getChannel().QueueDeclare(key, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	return q
}

func (c ClientImpl) Consume(key string, callback queue.CallbackFunc) {
	msgs, err := c.getChannel().Consume(
		key,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to consume a queue. Error: %s", err)
	}

	go func() {
		for d := range msgs {
			callback(d)
		}
	}()
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
