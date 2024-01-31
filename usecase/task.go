package usecase

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/kritmet/go-task/app"
	"github.com/kritmet/go-task/app/config"
	"github.com/kritmet/go-task/app/database"
	"github.com/kritmet/go-task/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type taskUsecase struct {
	returnResult config.ReturnResult
	repository   domain.TaskRepository
}

// NewTaskUsecase is a function for create task usecase
func NewTaskUsecase(application *app.Application, repository domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{
		repository:   repository,
		returnResult: application.ReturnResult,
	}
}

// Create is a function for create task
func (uc *taskUsecase) Create(c *gin.Context, req domain.CreateTaskRequest) error {
	var task domain.Task
	_ = copier.Copy(&task, &req)
	err := uc.repository.Create(database.Get(c), &task)
	if err != nil {
		logrus.Errorf("create error: %s", err)
		return err
	}

	return nil
}

// Update is a function for update task
func (uc *taskUsecase) Update(c *gin.Context, req domain.UpdateTaskRequest) error {
	task := domain.Task{}
	_ = copier.Copy(&task, &req)
	err := uc.repository.Update(database.Get(c), &task)
	if err != nil {
		logrus.Errorf("update error: %s", err)
		return err
	}

	return nil
}

// UpdateStatus is a function for update status task
func (uc *taskUsecase) UpdateStatus(c *gin.Context, req domain.UpdateStatusTaskRequest) error {
	err := uc.repository.UpdateFieldByID(database.Get(c), req.ID, "status", req.Status, &domain.Task{})
	if err != nil {
		logrus.Errorf("update status error: %s", err)
		return err
	}

	return nil
}

// Delete is a function for delete task
func (uc *taskUsecase) Delete(c *gin.Context, id uint) error {
	task, err := uc.repository.FindOneByID(database.Get(c), id)
	if err != nil {
		logrus.Errorf("find one by id error: %s", err)
		return err
	}

	err = uc.repository.Delete(database.Get(c), task)
	if err != nil {
		logrus.Errorf("delete error: %s", err)
		return err
	}

	return nil
}

// GetOneByID is a function for get one task by id
func (uc *taskUsecase) GetOneByID(c *gin.Context, id uint) (domain.Task, error) {
	task, err := uc.repository.FindOneByID(database.Get(c), id)
	if err != nil {
		logrus.Errorf("find one by id error: %s", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Task{}, uc.returnResult.DatabaseNotFound
		}
		return domain.Task{}, err
	}

	return task, nil
}

// GetAll is a function for get all task
func (uc *taskUsecase) GetAll(c *gin.Context, query domain.GetAllTaskRequest) ([]domain.Task, error) {
	tasks, err := uc.repository.FindAll(database.Get(c), query)
	if err != nil {
		logrus.Errorf("find all error: %s", err)
		return nil, err
	}

	return tasks, nil
}
