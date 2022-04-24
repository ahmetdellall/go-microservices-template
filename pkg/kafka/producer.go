package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"go-microservices-template/pkg/logger"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producer struct {
	log     logger.Logger
	brokers []string
	writer  *kafka.Writer
}

// create new kafka producer
func NewProducer(log logger.Logger, brokers []string) *producer {
	return &producer{
		log:     log,
		brokers: brokers,
		writer:  NewWriter(brokers, kafka.LoggerFunc(log.Errorf)),
	}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.writer.WriteMessages(ctx, msgs...)
}

func (p *producer) Close() error {
	return p.writer.Close()
}

// NewWriter create new configured kafka writer
func NewWriter(brokers []string, errLogger kafka.Logger) *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: writerRequiredAcks,
		MaxAttempts:  writerMaxAttempts,
		ErrorLogger:  errLogger,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
		Async:        false,
	}
	return w
}
