package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variables for MongoDB client and RabbitMQ channel
var (
	mongoClient  *mongo.Client
	rabbitMQConn *amqp.Connection
	rabbitMQCh   *amqp.Channel
)

// Data represents the MongoDB document structure
type Recipient struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MetaData struct {
	Priority   string `json:"priority"`
	Retries    int    `json:"retries"`
	ExecuteAt  int64  `json:"execute_at" bson:"execute_at"`
	MaxExecute int64  `json:"max_execute_at" bson:"max_execute_at"`
}

type RequestNotification struct {
	ID         string      `json:"id,omitempty" bson:"_id,omitempty"`
	Recipient  []Recipient `json:"recipient"`
	TemplateId string      `json:"template_id" bson:"template_id"`
	Quota      int64       `json:"quota"`
	Agent      string      `json:"agent"`
	MetaData   MetaData    `json:"meta_data" bson:"meta_data"`
	Status     string      `json:"status" bson:"status"`
}

// initMongoDB initializes a global MongoDB client
func initMongoDB(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetMaxPoolSize(50))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	return client
}

// initRabbitMQ initializes a global RabbitMQ connection and channel
func initRabbitMQ(uri string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ channel: %v", err)
	}

	return conn, ch
}

// queryMongoDB queries data from MongoDB and publishes it to RabbitMQ
func queryMongoDB(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query MongoDB
	cursor, err := collection.Find(ctx, bson.M{"status": "pending"})
	if err != nil {
		log.Printf("Failed to query MongoDB: %v", err)
		return
	}
	defer cursor.Close(ctx)

	// Process and publish each document
	for cursor.Next(ctx) {
		var data RequestNotification
		if err := cursor.Decode(&data); err != nil {
			log.Printf("Failed to decode MongoDB document: %v", err)
			continue
		}

		objectID, err := primitive.ObjectIDFromHex(data.ID)
		if err != nil {
			log.Fatalf("Invalid ObjectID: %v", err)
		}
		filter := bson.M{"_id": objectID}
		update := bson.M{
			"$set": bson.M{
				"status": "populated",
			},
		}
		res, err := collection.UpdateOne(ctx, filter, update)
		if err != nil && res.ModifiedCount > 0 {
			log.Printf("error Update Status: %v", err)
			continue
		}

		data.Status = "populated"
		message, err := json.Marshal(data)
		if err != nil {
			log.Printf("Failed to marshal data to JSON: %v", err)
			continue
		}

		// Publish to RabbitMQ
		err = rabbitMQCh.Publish(
			"notification_exchange",    // exchange
			"notification_routing_key", // routing key
			false,                      // mandatory
			false,                      // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        message,
			},
		)
		if err != nil {
			log.Printf("Failed to publish to RabbitMQ: %v", err)
		} else {
			log.Printf("Published message to RabbitMQ: %s", message)
		}
	}
}

func main() {
	// Initialize MongoDB and RabbitMQ
	mongoURI := "mongodb://root:asdf;lkj@localhost:27017"
	rabbitMQURI := "amqp://guest:guest@localhost:5672/"

	mongoClient = initMongoDB(mongoURI)
	defer mongoClient.Disconnect(context.Background())

	rabbitMQConn, rabbitMQCh = initRabbitMQ(rabbitMQURI)
	defer rabbitMQConn.Close()
	defer rabbitMQCh.Close()

	// Declare RabbitMQ queue
	err := rabbitMQCh.ExchangeDeclare(
		"notification_exchange", // Nama exchange
		"direct",                // Jenis exchange: "direct", "topic", "fanout", dll.
		true,                    // durable: exchange akan bertahan setelah restart RabbitMQ
		false,                   // auto-deleted: apakah exchange akan terhapus setelah tidak ada yang menggunakannya
		false,                   // internal: hanya bisa digunakan oleh RabbitMQ
		false,                   // no-wait: untuk menunggu feedback dari RabbitMQ
		nil,                     // Argumen tambahan
	)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	// MongoDB collection
	collection := mongoClient.Database("engine_notification").Collection("request_notifications")

	// Ticker to run the query every 5 minutes
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	// Run the query periodically
	for {
		select {
		case <-ticker.C:
			log.Println("Querying MongoDB and publishing to RabbitMQ...")
			queryMongoDB(collection)
		}
	}
}
