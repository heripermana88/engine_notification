package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func publishMessage(rabbitMQURL, exchangeName, routingKey, message string) error {
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

	// Mendeklarasikan exchange (jika belum ada)
	err = ch.ExchangeDeclare(
		exchangeName, // Nama exchange
		"direct",     // Jenis exchange: "direct", "topic", "fanout", dll.
		true,         // durable: exchange akan bertahan setelah restart RabbitMQ
		false,        // auto-deleted: apakah exchange akan terhapus setelah tidak ada yang menggunakannya
		false,        // internal: hanya bisa digunakan oleh RabbitMQ
		false,        // no-wait: untuk menunggu feedback dari RabbitMQ
		nil,          // Argumen tambahan
	)
	if err != nil {
		return err
	}

	// Mempublikasikan pesan
	err = ch.Publish(
		exchangeName, // Nama exchange
		routingKey,   // Routing key (dalam hal ini "permana")
		false,        // mandatory: apakah wajib ada consumer
		false,        // immediate: apakah pesan langsung diproses
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Pesan berhasil dipublikasikan dengan routing key %s", routingKey)
	return nil
}

func main() {
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"
	exchangeName := "my_exchange"
	routingKey := "permana"
	message := "Halo, ini pesan untuk key permana!"

	err := publishMessage(rabbitMQURL, exchangeName, routingKey, message)
	if err != nil {
		log.Fatalf("Gagal mengirim pesan: %s", err)
		os.Exit(1)
	}
}
