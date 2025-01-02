package adapter

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	channel *amqp.Channel
}

func NewRabbitMQPublisher(conn *amqp.Connection) (*RabbitMQPublisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{
		channel: ch,
	}, nil
}

func (p *RabbitMQPublisher) Publish(queueName string, message []byte) error {
	// Declare the queue if it doesn't already exist
	_, err := p.channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	// Publish the message
	err = p.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Message published to queue: %s", queueName)
	return nil
}

func (p *RabbitMQPublisher) Close() {
	p.channel.Close()
}
