package mappers

import (
	uuid "github.com/satori/go.uuid"
	"go-microservices-template/go_task_service/internal/models"
	writerService "go-microservices-template/go_task_service/proto"
	kafkaMessages "go-microservices-template/proto/kafka"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TaskToGrpcMessage(Task *models.Task) *kafkaMessages.Task {
	return &kafkaMessages.Task{
		TaskID:           Task.ID.String(),
		Title:            Task.Title,
		TaskKey:          Task.TaskKey,
		Details:          Task.Details,
		ExpectedDateTime: timestamppb.New(Task.ExpectedDateTime),
		Status:           int32(Task.Status),
		CreatorID:        Task.CreatorID.String(),
		CreatedDate:      timestamppb.New(Task.CreatedDate),
		UpdatedDate:      timestamppb.New(Task.UpdatedDate),
		StartDateTime:    timestamppb.New(Task.StartDateTime),
		EndDateTime:      timestamppb.New(Task.EndDateTime),
		CategoryID:       Task.CategoryID.String(),
	}
}

func TaskFromGrpcMessage(Task *kafkaMessages.Task) (*models.Task, error) {

	proUUID, err := uuid.FromString(Task.GetTaskID())
	if err != nil {
		return nil, err
	}

	creatorID, err := uuid.FromString(Task.GetCreatorID())
	if err != nil {
		return nil, err
	}

	categoryID, err := uuid.FromString(Task.GetCategoryID())
	if err != nil {
		return nil, err
	}

	return &models.Task{
		ID:               proUUID,
		Title:            Task.GetTitle(),
		TaskKey:          Task.GetTaskKey(),
		Details:          Task.GetDetails(),
		ExpectedDateTime: Task.GetExpectedDateTime().AsTime(),
		Status:           int(Task.GetStatus()),
		CreatorID:        creatorID,
		CreatedDate:      Task.CreatedDate.AsTime(),
		UpdatedDate:      Task.UpdatedDate.AsTime(),
		StartDateTime:    Task.StartDateTime.AsTime(),
		EndDateTime:      Task.EndDateTime.AsTime(),
		CategoryID:       categoryID,
	}, nil
}

func WriterTaskToGrpc(Task *models.Task) *writerService.Task {
	return &writerService.Task{
		TaskID:           Task.ID.String(),
		Title:            Task.Title,
		TaskKey:          Task.TaskKey,
		Details:          Task.Details,
		ExpectedDateTime: timestamppb.New(Task.ExpectedDateTime),
		Status:           int32(Task.Status),
		CreatorID:        Task.CreatorID.String(),
		CreatedDate:      timestamppb.New(Task.CreatedDate),
		UpdatedDate:      timestamppb.New(Task.UpdatedDate),
		StartDateTime:    timestamppb.New(Task.StartDateTime),
		EndDateTime:      timestamppb.New(Task.EndDateTime),
		CategoryID:       Task.CategoryID.String(),
	}
}
