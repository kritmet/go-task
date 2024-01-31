package repository

import (
	"fmt"

	"github.com/kritmet/go-task/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	Repository
}

// NewTaskRepository new repository
func NewTaskRepository() domain.TaskRepository {
	return &taskRepository{
		Repository: NewRepository(),
	}
}

// FindOneByID find one by id
func (r *taskRepository) FindOneByID(db *gorm.DB, id uint) (domain.Task, error) {
	var task domain.Task
	err := r.Repository.FindOneByField(db, "id", id, &task)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

// FindAll find all
func (r *taskRepository) FindAll(db *gorm.DB, query domain.GetAllTaskRequest) ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.query(db, query).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) query(db *gorm.DB, query domain.GetAllTaskRequest) *gorm.DB {
	if query.ID != 0 {
		return db.Where("id = ?", query.ID)
	}

	if query.Title != "" {
		return db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", query.Title))
	}

	if query.Description != "" {
		return db.Where("description LIKE ?", fmt.Sprintf("%%%s%%", query.Description))
	}

	if query.Status != "" {
		return db.Where("status = ?", query.Status)
	}

	return db
}
