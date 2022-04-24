package grpc

import (
	"context"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"go-microservices-template/go_task_service/config"
	"go-microservices-template/go_task_service/internal/handler"
	"go-microservices-template/go_task_service/internal/mappers"
	"go-microservices-template/go_task_service/internal/metrics"
	"go-microservices-template/go_task_service/internal/queries"
	"go-microservices-template/go_task_service/internal/service"
	writerService "go-microservices-template/go_task_service/proto"
	"go-microservices-template/pkg/logger"
	"go-microservices-template/pkg/tracing"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcService struct {
	log     logger.Logger
	cfg     *config.Config
	v       *validator.Validate
	ps      *service.TaskService
	metrics *metrics.WriterServiceMetrics
}

func NewWriterGrpcService(log logger.Logger, cfg *config.Config, v *validator.Validate, ps *service.TaskService, metrics *metrics.WriterServiceMetrics) *grpcService {
	return &grpcService{log: log, cfg: cfg, v: v, ps: ps, metrics: metrics}
}

func (s *grpcService) CreateTask(ctx context.Context, req *writerService.CreateTaskReq) (*writerService.CreateTaskRes, error) {
	s.metrics.CreateTaskGrpcRequests.Inc()

	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.CreateTask")
	defer span.Finish()

	TaskUUID, err := uuid.FromString(req.GetTaskID())
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	creatorId, err := uuid.FromString(req.CreatorID)
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	categoryId, err := uuid.FromString(req.CategoryID)
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	command := handler.NewCreateTaskCommand(TaskUUID, req.Title, req.TaskKey, req.Details, req.ExpectedDateTime.AsTime(), int(req.Status),
		creatorId, req.StartDateTime.AsTime(), req.EndDateTime.AsTime(), categoryId)
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	err = s.ps.Commands.CreateTask.Handle(ctx, command)
	if err != nil {
		s.log.WarnMsg("CreateTask.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	s.metrics.SuccessGrpcRequests.Inc()
	return &writerService.CreateTaskRes{TaskID: TaskUUID.String()}, nil
}

func (s *grpcService) UpdateTask(ctx context.Context, req *writerService.UpdateTaskReq) (*writerService.UpdateTaskRes, error) {
	s.metrics.UpdateTaskGrpcRequests.Inc()

	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.UpdateTask")
	defer span.Finish()

	TaskUUID, err := uuid.FromString(req.GetTaskID())
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	creatorId, err := uuid.FromString(req.CreatorID)
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	categoryId, err := uuid.FromString(req.CategoryID)
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	command := handler.NewUpdateTaskCommand(TaskUUID, req.Title, req.TaskKey, req.Details, req.ExpectedDateTime.AsTime(), int(req.Status),
		creatorId, req.StartDateTime.AsTime(), req.EndDateTime.AsTime(), categoryId)
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	err = s.ps.Commands.UpdateTask.Handle(ctx, command)
	if err != nil {
		s.log.WarnMsg("UpdateTask.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	s.metrics.SuccessGrpcRequests.Inc()
	return &writerService.UpdateTaskRes{}, nil
}

func (s *grpcService) GetTaskById(ctx context.Context, req *writerService.GetTaskByIdReq) (*writerService.GetTaskByIdRes, error) {
	s.metrics.GetTaskByIdGrpcRequests.Inc()

	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.GetTaskById")
	defer span.Finish()

	TaskUUID, err := uuid.FromString(req.GetTask().TaskID)
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	query := queries.NewGetTaskByIdQuery(TaskUUID)
	if err := s.v.StructCtx(ctx, query); err != nil {
		s.log.WarnMsg("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	Task, err := s.ps.Queries.GetTaskById.Handle(ctx, query)
	if err != nil {
		s.log.WarnMsg("GetTaskById.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	s.metrics.SuccessGrpcRequests.Inc()
	return &writerService.GetTaskByIdRes{Task: mappers.WriterTaskToGrpc(Task)}, nil
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	s.metrics.ErrorGrpcRequests.Inc()
	return status.Error(c, err.Error())
}
