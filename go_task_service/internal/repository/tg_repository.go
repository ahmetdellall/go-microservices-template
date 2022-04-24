package repository

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go-microservices-template/go_task_service/config"
	"go-microservices-template/go_task_service/internal/models"
	"go-microservices-template/pkg/logger"
	"gorm.io/gorm"
)

type taskRepository struct {
	log logger.Logger
	cfg *config.Config
	db  *gorm.DB
}

func NewTaskRepository(log logger.Logger, cfg *config.Config, db *gorm.DB) *taskRepository {
	return &taskRepository{log: log, cfg: cfg, db: db}
}

func (t *taskRepository) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "taskRepository.CreateTask")
	// wait func process finish
	defer span.Finish()

	var created models.Task

	if created := t.db.Create(&task); created.Error != nil {
		return nil, errors.Wrap(created.Error, "db.QueryRow")
	}
	return &created, nil
}

func (t *taskRepository) UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "taskRepository.UpdateTask")
	defer span.Finish()

	var update models.Task

	if update := t.db.Save(&task); update.Error != nil {
		return nil, errors.Wrap(update.Error, "Scan")
	}
	return &update, nil

}

func (t *taskRepository) GetTaskById(ctx context.Context, uuid uuid.UUID) (*models.Task, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "taskRepository.GetTaskById")
	defer span.Finish()
	var task models.Task

	if task := t.db.First(&uuid); task.Error != nil {
		return nil, errors.Wrap(task.Error, "Scan")
	}
	return &task, nil
}

func (t *taskRepository) DeleteTaskById(ctx context.Context, uuid uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "taskRepository.DeleteTaskByID")
	defer span.Finish()

	if tx := t.db.Delete(&models.Task{}, &uuid); tx.Error != nil {
		return errors.Wrap(tx.Error, "Exec")
	}

	return nil
}
