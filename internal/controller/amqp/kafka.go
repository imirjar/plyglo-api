package amqp

import (
	"github.com/segmentio/kafka-go"
)

type AMQP struct {
	writer  *kafka.Writer
	reader  *kafka.Reader
	service Service
}

type Service interface {
}

func New(opts ...func(*AMQP)) *AMQP {
	amqp := &AMQP{}
	for _, opt := range opts {
		opt(amqp)
	}

	return amqp
}

// writer := kafka.NewWriter(kafka.WriterConfig{})
//
//	defer writer.Close()
//	reader := kafka.NewReader(kafka.ReaderConfig{})
//	defer reader.Close()
func WithReader(service Service) func(*AMQP) {
	return func(amqp *AMQP) {
		amqp.service = service
	}
}
func WithWriter(service Service) func(*AMQP) {
	return func(amqp *AMQP) {
		amqp.service = service
	}
}
