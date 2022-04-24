package handler

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"go-microservices-template/go_task_service/config"
	"go-microservices-template/go_task_service/internal/mappers"
	"go-microservices-template/go_task_service/internal/models"
	"go-microservices-template/go_task_service/internal/repository"
	kafkaClient "go-microservices-template/pkg/kafka"
	"go-microservices-template/pkg/logger"
	"go-microservices-template/pkg/tracing"
	kafkaMessages "go-microservices-template/proto/kafka"
	"google.golang.org/protobuf/proto"
	"time"
)

type UpdateTaskCmdHandler interface {
	Handle(ctx context.Context, command *UpdateTaskCommand) error
}

type updateTaskHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer kafkaClient.Producer
}

func NewUpdateTaskHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *updateTaskHandler {
	return &updateTaskHandler{log: log, cfg: cfg, pgRepo: pgRepo, kafkaProducer: kafkaProducer}
}

func (c *updateTaskHandler) Handle(ctx context.Context, command *UpdateTaskCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateTaskHandler.Handle")
	defer span.Finish()

	TaskDto := &models.Task{
		ID:               command.TaskID,
		Title:            command.Title,
		TaskKey:          command.TaskKey,
		Details:          command.Details,
		ExpectedDateTime: command.ExpectedDateTime,
		Status:           command.Status,
		CreatorID:        command.CreatorID,
		StartDateTime:    command.StartDateTime,
		EndDateTime:      command.EndDateTime,
		CategoryID:       command.CategoryID,
	}

	Task, err := c.pgRepo.UpdateTask(ctx, TaskDto)
	if err != nil {
		return err
	}

	msg := &kafkaMessages.TaskUpdated{Task: mappers.TaskToGrpcMessage(Task)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.TaskUpdated.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
