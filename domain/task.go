package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TaskStatus task status
type TaskStatus string

const (
	// TaskStatusTodo task status todo
	TaskStatusTodo TaskStatus = "todo"
	// TaskStatusInProgress task status in progress
	TaskStatusInProgress TaskStatus = "in_progress"
	// TaskStatusDone task status done
	TaskStatusDone TaskStatus = "done"
)

// Task task model
type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}

// CreateTaskRequest create task request
type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Status      TaskStatus `json:"status" binding:"required,oneof=todo in_progress done"`
}

// UpdateTaskRequest update task request
type UpdateTaskRequest struct {
	ID uint `json:"-"`
	CreateTaskRequest
}

// UpdateStatusTaskRequest update status task request
type UpdateStatusTaskRequest struct {
	ID     uint       `json:"-"`
	Status TaskStatus `json:"status" binding:"required,oneof=todo in_progress done"`
}

// GetAllTaskRequest get all task request
type GetAllTaskRequest struct {
	ID          uint       `form:"id"`
	Title       string     `form:"title"`
	Description string     `form:"description"`
	Status      TaskStatus `form:"status"`
}

// TaskController task controller
type TaskController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetOneByID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

// TaskUsecase task usecase
type TaskUsecase interface {
	Create(c *gin.Context, request CreateTaskRequest) error
	Update(c *gin.Context, request UpdateTaskRequest) error
	UpdateStatus(c *gin.Context, request UpdateStatusTaskRequest) error
	Delete(c *gin.Context, id uint) error
	GetOneByID(c *gin.Context, id uint) (Task, error)
	GetAll(c *gin.Context, query GetAllTaskRequest) ([]Task, error)
}

// TaskRepository task repository
type TaskRepository interface {
	Create(db *gorm.DB, i interface{}) error
	Update(db *gorm.DB, i interface{}) error
	UpdateFieldByID(db *gorm.DB, id uint, field string, value, i interface{}) error
	Delete(db *gorm.DB, i interface{}) error
	FindOneByID(db *gorm.DB, id uint) (Task, error)
	FindAll(db *gorm.DB, query GetAllTaskRequest) ([]Task, error)
}
