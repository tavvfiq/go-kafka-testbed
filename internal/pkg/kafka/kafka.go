package kafka

import (
	"context"
	"time"

	kfk "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

type KafkaWriter interface {
	Push(ctx context.Context, message kfk.Message) error
}

type KafkaReader interface {
}

type service struct {
	writer *kfk.Writer
	// reader *kfk.Reader
}

func NewWriter(urls []string, clientId, topic string, timeout time.Duration) KafkaWriter {
	dialer := &kfk.Dialer{
		Timeout:  timeout,
		ClientID: clientId,
	}
	config := kfk.WriterConfig{
		Brokers:          urls,
		Topic:            topic,
		Balancer:         &kfk.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     timeout,
		ReadTimeout:      timeout,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	return &service{
		writer: kfk.NewWriter(config),
	}
}

func NewReader(config kfk.ReaderConfig) KafkaReader {
	return service{}
}
