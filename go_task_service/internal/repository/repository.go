package repository

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go-microservices-template/go_task_service/internal/models"
)

type Repository interface {
	CreateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	DeleteTaskByID(ctx context.Context, uuid uuid.UUID) error

	GetTaskById(ctx context.Context, uuid uuid.UUID) (*models.Task, error)
}
