package adapter

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	channel  *amqp.Channel
	exchange string
}

func NewRabbitMQPublisher(conn *amqp.Connection, exchange string) (*RabbitMQPublisher, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Declare an exchange
	err = channel.ExchangeDeclare(
		exchange, // exchange name
		"direct", // exchange type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{channel: channel, exchange: exchange}, nil
}

func (p *RabbitMQPublisher) PublishMessage(routingKey string, body []byte) error {
	return p.channel.Publish(
		p.exchange, // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

type RabbitMQListener struct {
	channel    *amqp.Channel
	exchange   string
	queue      string
	routingKey string
	callback   func([]byte) error
}

func NewRabbitMQListener(conn *amqp.Connection, exchange, routingKey string, callback func([]byte) error) (*RabbitMQListener, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Declare a queue
	queue, err := channel.QueueDeclare(
		"",    // queue name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	// Bind the queue to the exchange with the specified routing key
	err = channel.QueueBind(
		queue.Name, // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQListener{
		channel:    channel,
		exchange:   exchange,
		queue:      queue.Name,
		routingKey: routingKey,
		callback:   callback,
	}, nil
}

func (l *RabbitMQListener) StartListening() error {
	msgs, err := l.channel.Consume(
		l.queue, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		err := l.callback(msg.Body)
		if err != nil {
			log.Printf("Failed to process message: %v", err)
		}
	}

	return nil
}
