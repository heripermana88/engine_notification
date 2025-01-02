package adapter

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQListener struct {
	channel  *amqp.Channel
	queue    string
	callback func([]byte) error
}

func NewRabbitMQListener(conn *amqp.Connection, queueName string, callback func([]byte) error) (*RabbitMQListener, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Declare the queue
	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQListener{
		channel:  ch,
		queue:    queueName,
		callback: callback,
	}, nil
}

func (l *RabbitMQListener) StartListening() error {
	msgs, err := l.channel.Consume(
		l.queue,
		"",    // consumer tag
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received message listener: %s", d.Body)
			if err := l.callback(d.Body); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}()

	return nil
}

func (l *RabbitMQListener) Close() {
	l.channel.Close()
}
