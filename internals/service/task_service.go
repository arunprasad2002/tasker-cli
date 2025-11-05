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

func (s *TaskService) ListTasks(status models.TaskStatus) ([]models.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// If no status is provided, return all tasks
	if status == "" {
		return tasks, nil
	}

	var filteredTasks []models.Task

	for _, t := range tasks {
		if t.Status == status {
			filteredTasks = append(filteredTasks, t)
		}
	}

	return filteredTasks, nil
}

func (s *TaskService) MarkStatus(id uint, status models.TaskStatus) (string, error) {
	return s.repo.UpdateTask(id, "", status)
}

func (s *TaskService) DeleteTask(id uint) (string, error) {
	return s.repo.Delete(id)
}

func (s *TaskService) UpdateTask(id uint, title string, status models.TaskStatus) (string, error) {
	return s.repo.UpdateTask(id, title, status)
}
