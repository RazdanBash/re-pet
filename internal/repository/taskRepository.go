package repository

import (
	"gorm.io/gorm"
	"re-pet/internal/models"
)

type TaskRepositoryDB struct {
	db *gorm.DB
}

type TaskRepository interface {
	getAll() ([]models.Task, error)
	Create(task models.Task) (models.Task, error)
	Update(id uint, task models.Task) (models.Task, error)
	Delete(id uint) error
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryDB {
	return &TaskRepositoryDB{db: db}
}

func (r *TaskRepositoryDB) getAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
