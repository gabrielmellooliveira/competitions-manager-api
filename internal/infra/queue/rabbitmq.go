package queue

import (
	"log"

	queue "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/queue"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMqAdapter struct {
	Url string
}

func NewRabbitMqAdapter(url string) queue.Queue {
	return &RabbitMqAdapter{
		Url: url,
	}
}

func (e *RabbitMqAdapter) Consume(queueName string, handler queue.Handler) {
	connection := e.getConnection()
	defer connection.Close()

	channel := e.getChannel(connection)
	defer channel.Close()

	_, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("falha ao declarar a fila: %v", err)
	}

	delivery := e.getDelivery(channel, queueName)

	done := make(chan bool)

	go func() {
		for message := range delivery {
			err := handler(message.Body)
			if err != nil {
				println(err.Error())
			}
		}
	}()

	<-done
}

func (e *RabbitMqAdapter) Publish(queueName string, message []byte) error {
	connection := e.getConnection()
	defer connection.Close()

	channel := e.getChannel(connection)
	defer channel.Close()

	_, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("falha ao declarar a fila: %v", err)
		return err
	}

	err = channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		log.Fatalf("falha ao publicar a mensagem: %v", err)
		return err
	}

	return nil
}

func (e *RabbitMqAdapter) getConnection() *amqp091.Connection {
	conn, err := amqp091.Dial(e.Url)
	if err != nil {
		log.Fatalf("falha ao conectar ao RabbitMQ: %v", err)
	}

	return conn
}

func (e *RabbitMqAdapter) getChannel(connection *amqp091.Connection) *amqp091.Channel {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("falha ao abrir um novo canal: %v", err)
	}

	return ch
}

func (e *RabbitMqAdapter) getDelivery(channel *amqp091.Channel, queueName string) <-chan amqp091.Delivery {
	msgs, err := channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("falha ao registrar um consumidor: %v", err)
	}

	return msgs
}
