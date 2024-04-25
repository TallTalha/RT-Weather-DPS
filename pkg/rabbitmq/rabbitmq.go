package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQClient RabbitMQ ile bağlantıyı tutacak yapı
type RabbitMQClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbitMQClient yeni bir RabbitMQ bağlantısı oluşturur ve döndürür
func NewRabbitMQClient(url string) *RabbitMQClient {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	return &RabbitMQClient{
		Connection: conn,
		Channel:    ch,
	}
}

// PublishMessage belirtilen kuyruk adına mesaj gönderir
func (client *RabbitMQClient) PublishMessage(queueName string, body []byte) {
	_, err := client.Channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = client.Channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
}

// Close bağlantıyı ve kanalı kapatır
func (client *RabbitMQClient) Close() {
	client.Channel.Close()
	client.Connection.Close()
}
