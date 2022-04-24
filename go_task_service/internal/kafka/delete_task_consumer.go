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

func (s *TaskMessageProcessor) processDeleteTask(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.DeleteTaskKafkaMessages.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "TaskMessageProcessor.processDeleteTask")
	defer span.Finish()

	msg := &kafkaMessages.TaskDelete{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	proUUID, err := uuid.FromString(msg.GetTaskId())
	if err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	command := handler.NewDeleteTaskCommand(proUUID)
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return s.ps.Commands.DeleteTask.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("DeleteTask.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}

	s.commitMessage(ctx, r, m)
}
