package services

import (
	"re-pet/internal/models"
	"re-pet/internal/repository"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
}

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	return s.repo.Create(task)
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) UpdateTask(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
