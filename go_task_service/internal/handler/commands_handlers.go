package handler

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type TaskCommands struct {
	CreateTask createTaskCmdHandler
	UpdateTask UpdateTaskCmdHandler
	DeleteTask DeleteTaskCmdHandler
}

func NewTaskCommands(createTask createTaskCmdHandler, updateTask UpdateTaskCmdHandler, deleteTask DeleteTaskCmdHandler) *TaskCommands {
	return &TaskCommands{CreateTask: createTask, UpdateTask: updateTask, DeleteTask: deleteTask}
}

type CreateTaskCommand struct {
	TaskID           uuid.UUID `json:"TaskId" validate:"required"`
	Title            string    `json:"title" validate:"required,gte=0,lte=255"`
	TaskKey          string    `json:"taskKey" validate:"required,gte=0,lte=5000"`
	Details          string    `json:"details" validate:"required,gte=0,lte=5000"`
	ExpectedDateTime time.Time `json:"expectedDateTime"`
	Status           int       `json:"status" validate:"required,gte=0""`
	CreatorID        uuid.UUID `json:"creatorID" validate:"required"`
	StartDateTime    time.Time `json:"startDateTime"`
	EndDateTime      time.Time `json:"endDateTime"`
	CategoryID       uuid.UUID `json:"categoryID" validate:"required"`
}

func NewCreateTaskCommand(TaskID uuid.UUID, Title string, TaskKey string, Details string, ExpectedDateTime time.Time, Status int, CreatorID uuid.UUID, StartDateTime time.Time, EndDateTime time.Time, CategoryID uuid.UUID) *CreateTaskCommand {
	return &CreateTaskCommand{
		TaskID:           TaskID,
		Title:            Title,
		TaskKey:          TaskKey,
		Details:          Details,
		ExpectedDateTime: ExpectedDateTime,
		Status:           Status,
		CreatorID:        CreatorID,
		StartDateTime:    StartDateTime,
		EndDateTime:      EndDateTime,
		CategoryID:       CategoryID,
	}
}

type UpdateTaskCommand struct {
	TaskID           uuid.UUID `json:"taskID" validate:"required"`
	Title            string    `json:"title" validate:"required,gte=0,lte=255"`
	TaskKey          string    `json:"taskKey" validate:"required,gte=0,lte=5000"`
	Details          string    `json:"details" validate:"required,gte=0,lte=5000"`
	ExpectedDateTime time.Time `json:"expectedDateTime"`
	Status           int       `json:"status" validate:"required,gte=0""`
	CreatorID        uuid.UUID `json:"creatorID" validate:"required"`
	StartDateTime    time.Time `json:"startDateTime"`
	EndDateTime      time.Time `json:"endDateTime"`
	CategoryID       uuid.UUID `json:"categoryID" validate:"required"`
}

func NewUpdateTaskCommand(TaskID uuid.UUID, Title string, TaskKey string, Details string, ExpectedDateTime time.Time, Status int, CreatorID uuid.UUID, StartDateTime time.Time, EndDateTime time.Time, CategoryID uuid.UUID) *UpdateTaskCommand {
	return &UpdateTaskCommand{
		TaskID:           TaskID,
		Title:            Title,
		TaskKey:          TaskKey,
		Details:          Details,
		ExpectedDateTime: ExpectedDateTime,
		Status:           Status,
		CreatorID:        CreatorID,
		StartDateTime:    StartDateTime,
		EndDateTime:      EndDateTime,
		CategoryID:       CategoryID,
	}
}

type DeleteTaskCommand struct {
	TaskID uuid.UUID `json:"TaskId" validate:"required"`
}

func NewDeleteTaskCommand(TaskID uuid.UUID) *DeleteTaskCommand {
	return &DeleteTaskCommand{TaskID: TaskID}
}
