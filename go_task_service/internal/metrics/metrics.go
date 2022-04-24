package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go-microservices-template/go_task_service/config"
)

type WriterServiceMetrics struct {
	SuccessGrpcRequests prometheus.Counter
	ErrorGrpcRequests   prometheus.Counter

	CreateTaskGrpcRequests  prometheus.Counter
	UpdateTaskGrpcRequests  prometheus.Counter
	DeleteTaskGrpcRequests  prometheus.Counter
	GetTaskByIdGrpcRequests prometheus.Counter
	SearchTaskGrpcRequests  prometheus.Counter

	SuccessKafkaMessages prometheus.Counter
	ErrorKafkaMessages   prometheus.Counter

	CreateTaskKafkaMessages prometheus.Counter
	UpdateTaskKafkaMessages prometheus.Counter
	DeleteTaskKafkaMessages prometheus.Counter
}

func NewWriterServiceMetrics(cfg *config.Config) *WriterServiceMetrics {
	return &WriterServiceMetrics{
		SuccessGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of success grpc requests",
		}),
		ErrorGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of error grpc requests",
		}),
		CreateTaskGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_Task_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of create Task grpc requests",
		}),
		UpdateTaskGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_Task_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of update Task grpc requests",
		}),
		DeleteTaskGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_Task_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of delete Task grpc requests",
		}),
		GetTaskByIdGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_Task_by_id_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of get Task by id grpc requests",
		}),
		SearchTaskGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_Task_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of search Task grpc requests",
		}),
		CreateTaskKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_Task_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of create Task kafka messages",
		}),
		UpdateTaskKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_Task_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of update Task kafka messages",
		}),
		DeleteTaskKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_Task_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of delete Task kafka messages",
		}),
		SuccessKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of success kafka processed messages",
		}),
		ErrorKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of error kafka processed messages",
		}),
	}
}
