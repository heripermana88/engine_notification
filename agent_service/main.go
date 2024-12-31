package main

import (
	"log"

	"github.com/streadway/amqp"
)

func consumeMessages(rabbitMQURL, exchangeName, routingKey string) error {
	// Membuka koneksi ke RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Membuka channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Mendeklarasikan exchange (harus sama dengan yang digunakan untuk publish)
	err = ch.ExchangeDeclare(
		exchangeName, // Nama exchange
		"direct",     // Tipe exchange (sama dengan publisher)
		true,         // Durable
		false,        // Auto-deleted
		false,        // Internal
		false,        // No-wait
		nil,          // Argumen tambahan
	)
	if err != nil {
		return err
	}

	// Mendeklarasikan antrian untuk menerima pesan
	q, err := ch.QueueDeclare(
		"",    // Nama queue otomatis
		true,  // Durable
		false, // Auto-delete
		true,  // Exclusive
		false, // No-wait
		nil,   // Argumen tambahan
	)
	if err != nil {
		return err
	}

	// Binding queue ke exchange dengan routing key yang sama
	err = ch.QueueBind(
		q.Name,       // Nama queue
		routingKey,   // Routing key
		exchangeName, // Nama exchange
		false,        // No-wait
		nil,          // Argumen tambahan
	)
	if err != nil {
		return err
	}

	// Menunggu dan menerima pesan
	msgs, err := ch.Consume(
		q.Name, // Nama queue
		"",     // Consumer tag (biarkan kosong)
		true,   // Auto-acknowledge
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Argumen tambahan
	)
	if err != nil {
		return err
	}

	log.Printf("Menunggu pesan di queue %s", q.Name)

	// Menerima dan memproses pesan
	for msg := range msgs {
		log.Printf("Pesan diterima: %s", msg.Body)
	}

	return nil
}

func main() {
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"
	exchangeName := "notification_exchange"
	routingKey := "notification_routing_key"

	err := consumeMessages(rabbitMQURL, exchangeName, routingKey)
	if err != nil {
		log.Fatalf("Gagal menerima pesan: %s", err)
	}
}
