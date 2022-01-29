package service

import (
	"context"

	pb "home_work/task-service/genproto/task"
	l "home_work/task-service/pkg/logger"
	"home_work/task-service/storage"

	"github.com/jmoiron/sqlx"
)

//TaskService ...
type TaskService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewTaskService ...
func NewTaskService(db *sqlx.DB, log l.Logger) *TaskService {
	return &TaskService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

// Create Task ...
func (s *TaskService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	task, err := s.storage.Task().Create(req)
	if err != nil {
		s.logger.Error("Error create task", l.Error(err))
		return nil, nil
	}
	return task, nil
}

// Get Task ...
func (s *TaskService) Get(ctx context.Context, req *pb.IdReq) (*pb.Task, error) {
	task, err := s.storage.Task().Get(req.Id)
	if err != nil {
		s.logger.Error("Error get task", l.Error(err))
		return nil, nil
	}
	return task, nil
}

// List Tasks ...
func (s *TaskService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
    resp, err := s.storage.Task().List(req)
    if err != nil {
        s.logger.Error("Error list tasks", l.Error(err))
        return nil, nil
    }
	return resp, nil
}

// Update task
func (s *TaskService) Update(ctx context.Context, req *pb.Task) (*pb.Task, error) {
    task, err := s.storage.Task().Update(req)
    if err != nil {
        s.logger.Error("Error update task")
        return nil, nil
    }
	return task, nil
}

// Delete task
func (s *TaskService) Delete(ctx context.Context, req *pb.IdReq) (*pb.EmptyResp, error) {
    _, err := s.storage.Task().Delete(req)
    if err != nil {
        s.logger.Error("Error delete task")
        return nil, nil
    }
	return &pb.EmptyResp{}, nil
}

// ListOverdue
func (s *TaskService) ListOverdue(ctx context.Context, req *pb.ListOverReq) (*pb.ListOverResp, error) {
	return nil, nil
}
