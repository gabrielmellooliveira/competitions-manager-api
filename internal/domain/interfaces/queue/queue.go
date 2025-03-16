package interfaces

type Queue interface {
	Publish(queueName string, message []byte) error
	Consume(queueName string, handler Handler)
}

type Handler func(body []byte) error
