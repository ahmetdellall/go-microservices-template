package queries

import uuid "github.com/satori/go.uuid"

type TaskQueries struct {
	GetTaskById GetTaskByIdHandler
}

func NewTaskQueries(getTaskById GetTaskByIdHandler) *TaskQueries {
	return &TaskQueries{GetTaskById: getTaskById}
}

type GetTaskByIdQuery struct {
	TaskID uuid.UUID `json:"TaskId" validate:"required,gte=0,lte=255"`
}

func NewGetTaskByIdQuery(TaskID uuid.UUID) *GetTaskByIdQuery {
	return &GetTaskByIdQuery{TaskID: TaskID}
}
