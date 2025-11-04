package repository

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (r *TaskRepository) Delete(id uint) (string, error) {
	// step 1: read all tasks
	tasks, err := r.readTasks()

	if err != nil {
		return "", err
	}
	//step 2: create a new slice excluding the task with matching ID
	updatedTasks := []models.Task{}
	found := false

	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}

		updatedTasks = append(updatedTasks, task)
	}

	// Step 3: Handle case where task not found

	if !found {
		return "", fmt.Errorf("task with id %d not found", id)
	}

	err = r.writeTask(updatedTasks)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Task with ID %d deleted successfully.", id), nil
}

// Update Task

func (r *TaskRepository) UpdateTask(id uint, title string, status models.TaskStatus) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Step 1: Read all tasks
	tasks, err := r.readTasks()
	if err != nil {
		return "", err
	}

	found := false

	// Step 2: Update the matching task
	for i, task := range tasks {
		if task.ID == id {
			found = true

			if title != "" {
				tasks[i].Title = title
			}
			if status != "" {
				tasks[i].Status = status
			}

			tasks[i].UpdatedAt = time.Now()
			break
		}
	}

	// Step 3: Handle not found
	if !found {
		return "", fmt.Errorf("task with id %d not found", id)
	}

	// Step 4: Write the updated slice back to file
	if err := r.writeTask(tasks); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task with ID %d updated successfully.", id), nil
}

// Get All Tasks
func (r *TaskRepository) GetAll() ([]models.Task, error) {
	return r.readTasks()
}
