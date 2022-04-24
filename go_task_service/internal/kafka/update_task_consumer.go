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
)

func (s *TaskMessageProcessor) processUpdateTask(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.UpdateTaskKafkaMessages.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "TaskMessageProcessor.processUpdateTask")
	defer span.Finish()

	msg := &kafkaMessages.TaskUpdate{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
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

	command := handler.NewUpdateTaskCommand(proUUID, msg.Title, msg.TaskKey, msg.Details, msg.ExpectedDateTime.AsTime(), int(msg.Status),
		creatorId, msg.StartDateTime.AsTime(), msg.EndDateTime.AsTime(), categoryId)
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return s.ps.Commands.UpdateTask.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("UpdateTask.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}

	s.commitMessage(ctx, r, m)
}
