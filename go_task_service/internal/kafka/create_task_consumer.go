package kafka

import (
	"context"
	"github.com/avast/retry-go"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
	"go-microservices-template/go_task_service/internal/handler"
	"go-microservices-template/pkg/tracing"
	kafkaMessages "go-microservices-template/proto/kafka"
	"google.golang.org/protobuf/proto"
	"time"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (s *TaskMessageProcessor) processCreateTask(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.CreateTaskKafkaMessages.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "TaskMessageProcessor.processCreateTask")
	defer span.Finish()

	var msg kafkaMessages.TaskCreate
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	proUUID, err := uuid.FromString(msg.GetTaskID())
	if err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	creatorId, err := uuid.FromString(msg.CreatorID)
	if err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	categoryId, err := uuid.FromString(msg.CategoryID)
	if err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	command := handler.NewCreateTaskCommand(proUUID, msg.Title, msg.TaskKey, msg.Details, msg.ExpectedDateTime.AsTime(), int(msg.Status),
		creatorId, msg.StartDateTime.AsTime(), msg.EndDateTime.AsTime(), categoryId)
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return s.ps.Commands.CreateTask.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("CreateTask.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}

	s.commitMessage(ctx, r, m)
}
