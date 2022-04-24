package handler

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"go-microservices-template/go_task_service/config"
	"go-microservices-template/go_task_service/internal/repository"
	kafkaClient "go-microservices-template/pkg/kafka"
	"go-microservices-template/pkg/logger"
	"go-microservices-template/pkg/tracing"
	kafkaMessages "go-microservices-template/proto/kafka"
	"google.golang.org/protobuf/proto"
	"time"
)

type DeleteTaskCmdHandler interface {
	Handle(ctx context.Context, command *DeleteTaskCommand) error
}

type deleteTaskHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer kafkaClient.Producer
}

func NewDeleteTaskHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *deleteTaskHandler {
	return &deleteTaskHandler{log: log, cfg: cfg, pgRepo: pgRepo, kafkaProducer: kafkaProducer}
}

func (c *deleteTaskHandler) Handle(ctx context.Context, command *DeleteTaskCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteTaskHandler.Handle")
	defer span.Finish()

	if err := c.pgRepo.DeleteTaskByID(ctx, command.TaskID); err != nil {
		return err
	}

	msg := &kafkaMessages.TaskDeleted{TaskId: command.TaskID.String()}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.TaskDeleted.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
