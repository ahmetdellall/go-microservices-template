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

type createTaskCmdHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer kafkaClient.Producer
}

func NewCreateTaskHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *createTaskCmdHandler {
	return &createTaskCmdHandler{log: log, cfg: cfg, pgRepo: pgRepo, kafkaProducer: kafkaProducer}
}

func (c *createTaskCmdHandler) Handle(ctx context.Context, cmd *CreateTaskCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createTaskHandler.Handle")
	defer span.Finish()

	TaskDto := &models.Task{
		ID:               cmd.TaskID,
		Title:            cmd.Title,
		TaskKey:          cmd.TaskKey,
		Details:          cmd.Details,
		ExpectedDateTime: cmd.ExpectedDateTime,
		Status:           cmd.Status,
		CreatorID:        cmd.CreatorID,
		StartDateTime:    cmd.StartDateTime,
		EndDateTime:      cmd.EndDateTime,
		CategoryID:       cmd.CategoryID,
	}

	Task, err := c.pgRepo.CreateTask(ctx, TaskDto)
	if err != nil {
		return err
	}

	msg := &kafkaMessages.TaskCreated{Task: mappers.TaskToGrpcMessage(Task)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.TaskCreated.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
