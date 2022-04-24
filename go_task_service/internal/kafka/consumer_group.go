package kafka

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/segmentio/kafka-go"
	"go-microservices-template/go_task_service/config"
	"go-microservices-template/go_task_service/internal/metrics"
	"go-microservices-template/go_task_service/internal/service"
	"go-microservices-template/pkg/logger"
	"sync"
)

const (
	PoolSize = 30
)

type TaskMessageProcessor struct {
	log     logger.Logger
	cfg     *config.Config
	v       *validator.Validate
	ps      *service.TaskService
	metrics *metrics.WriterServiceMetrics
}

func NewTaskMessageProcessor(log logger.Logger, cfg *config.Config, v *validator.Validate, ps *service.TaskService, metrics *metrics.WriterServiceMetrics) *TaskMessageProcessor {
	return &TaskMessageProcessor{log: log, cfg: cfg, v: v, ps: ps, metrics: metrics}
}

func (s *TaskMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}

		s.logProcessMessage(m, workerID)

		switch m.Topic {
		case s.cfg.KafkaTopics.TaskCreate.TopicName:
			s.processCreateTask(ctx, r, m)
		case s.cfg.KafkaTopics.TaskUpdate.TopicName:
			s.processUpdateTask(ctx, r, m)
		case s.cfg.KafkaTopics.TaskDelete.TopicName:
			s.processDeleteTask(ctx, r, m)
		}
	}
}
