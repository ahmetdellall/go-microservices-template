package service

import (
	"go-microservices-template/go_task_service/config"
	"go-microservices-template/go_task_service/internal/handler"
	"go-microservices-template/go_task_service/internal/queries"
	"go-microservices-template/go_task_service/internal/repository"
	kafkaClient "go-microservices-template/pkg/kafka"
	"go-microservices-template/pkg/logger"
)

type TaskService struct {
	Commands *handler.TaskCommands
	Queries  *queries.TaskQueries
}

func NewTaskService(log logger.Logger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *TaskService {

	updateTaskHandler := handler.NewUpdateTaskHandler(log, cfg, pgRepo, kafkaProducer)
	createTaskHandler := handler.NewCreateTaskHandler(log, cfg, pgRepo, kafkaProducer)
	deleteTaskHandler := handler.NewDeleteTaskHandler(log, cfg, pgRepo, kafkaProducer)

	getTaskByIdHandler := queries.NewGetTaskByIdHandler(log, cfg, pgRepo)

	TaskCommands := handler.NewTaskCommands(createTaskHandler, updateTaskHandler, deleteTaskHandler)
	TaskQueries := queries.NewTaskQueries(getTaskByIdHandler)

	return &TaskService{Commands: TaskCommands, Queries: TaskQueries}
}
