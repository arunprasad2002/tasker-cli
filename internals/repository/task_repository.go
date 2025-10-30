package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"tasker/internals/models"
	"time"
)

type TaskRepository struct {
	filePath string
	mu       sync.Mutex
}

func NewTaskRepository(filePath string) *TaskRepository {
	return &TaskRepository{
		filePath: filePath,
	}
}

// helper: write task to file

func (r *TaskRepository) writeTask(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	// ✅ Ensure folder exists before writing
	if err := os.MkdirAll(filepath.Dir(r.filePath), 0755); err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

// helper: read tasks from file

func (r *TaskRepository) readTasks() ([]models.Task, error) {
	file, error := os.ReadFile(r.filePath)
	if errors.Is(error, os.ErrNotExist) {
		return []models.Task{}, nil //no file yet
	} else if error != nil {
		return nil, error
	}

	// ✅ FIX: handle empty files gracefully
	if len(file) == 0 {
		return []models.Task{}, nil
	}

	var tasks []models.Task

	err := json.Unmarshal(file, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Create Tasks

func (r *TaskRepository) Create(title string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.readTasks()

	if err != nil {
		return err
	}

	newTask := models.Task{
		ID:        uint(len(tasks) + 1),
		Title:     title,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)

	return r.writeTask(tasks)
}

// Get All Tasks
func (r *TaskRepository) GetAll() ([]models.Task, error) {
	return r.readTasks()
}
