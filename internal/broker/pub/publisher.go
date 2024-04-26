package pub

import (
	"log"

	"github.com/streadway/amqp"
)

// Publisher RabbitMQ ile mesaj yayınlama işlemlerini yönetir
type Publisher struct {
	Conn *amqp.Connection
}

// NewPublisher yeni bir Publisher örneği oluşturur
func NewPublisher(conn *amqp.Connection) *Publisher {
	return &Publisher{Conn: conn}
}

// PublishMessage belirtilen kuyruk adına mesaj gönderir
func (p *Publisher) PublishMessage(queueName string, body []byte) error {
	ch, err := p.Conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return err
	}

	err = ch.Publish(
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
		log.Printf("Failed to publish a message: %v", err)
		return err
	}
	log.Printf("Message published to queue %s", queueName)
	return nil
}
