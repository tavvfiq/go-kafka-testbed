package kafka

import (
	"context"

	kfk "github.com/segmentio/kafka-go"
)

func (s *service) Push(ctx context.Context, message kfk.Message) error {
	return s.writer.WriteMessages(ctx, message)
}
