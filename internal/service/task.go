package service

import (
	"context"

	"github.com/PainCodermax/to-do-list-api/internal/domain"
	"github.com/PainCodermax/to-do-list-api/internal/entity"
	"github.com/PainCodermax/to-do-list-api/internal/service/serializers"
	"github.com/google/uuid"
)

type TaskService struct {
	ctx        context.Context
	taskDomain *domain.TaskDomain
}

func NewTaskService(ctx context.Context) *TaskService {
	return &TaskService{
		ctx:        ctx,
		taskDomain: domain.NewTaskDomain()}
}

func (s *TaskService) Create(requestBody *serializers.TaskCreateRequest) (*entity.Task, error) {
	id := uuid.New().String()
	task, err := s.taskDomain.Create(id, requestBody.Description, requestBody.Title)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) GetAll() ([]entity.Task, error) {
	tasks := s.taskDomain.GetAll()
	return tasks, nil
}

func (s *TaskService) GetByID(id string) (*entity.Task, error) {
	task, err := s.taskDomain.GetByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Complete(id string) error {
	if _, err := s.taskDomain.Complete(id); err != nil {
		return err
	}
	return nil
}

func (s *TaskService) Delete(id string) error {
	if err := s.taskDomain.Delete(id); err != nil {
		return err
	}
	return nil
}
