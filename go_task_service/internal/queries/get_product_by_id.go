package queries

import (
	"context"
	"go-microservices-template/go_task_service/config"
	"go-microservices-template/go_task_service/internal/models"
	"go-microservices-template/go_task_service/internal/repository"
	"go-microservices-template/pkg/logger"
)

type GetTaskByIdHandler interface {
	Handle(ctx context.Context, query *GetTaskByIdQuery) (*models.Task, error)
}

type getTaskByIdHandler struct {
	log    logger.Logger
	cfg    *config.Config
	pgRepo repository.Repository
}

func NewGetTaskByIdHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository) *getTaskByIdHandler {
	return &getTaskByIdHandler{log: log, cfg: cfg, pgRepo: pgRepo}
}

func (q *getTaskByIdHandler) Handle(ctx context.Context, query *GetTaskByIdQuery) (*models.Task, error) {
	return q.pgRepo.GetTaskById(ctx, query.TaskID)
}
