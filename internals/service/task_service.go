package service

import (
	"tasker/internals/models"
	"tasker/internals/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(title string) error {
	return s.repo.Create(title)
}

func (s *TaskService) ListTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) DeleteTask(id uint) (string, error) {
	return s.repo.Delete(id)
}

func (s *TaskService) UpdateTask(id uint, title string, status models.TaskStatus) (string, error) {
	return s.repo.UpdateTask(id, title, status)
}
