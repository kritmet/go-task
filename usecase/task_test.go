package usecase

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-task/app"
	"github.com/kritmet/go-task/app/config"
	"github.com/kritmet/go-task/domain"
	"github.com/kritmet/go-task/repository/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func init() {
	err := config.InitReturnResult("../configs")
	if err != nil {
		logrus.Panic(err)
	}
}

type TestSuite struct {
	usecase        domain.TaskUsecase
	mockRepository *mocks.TaskRepository
	ctx            *gin.Context
	T              *testing.T
}

func (s *TestSuite) SetupTest() {
	s.mockRepository = mocks.NewTaskRepository(s.T)
	s.usecase = NewTaskUsecase(&app.Application{ReturnResult: config.RR}, s.mockRepository)
	s.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())

}

func TestCreateTask(t *testing.T) {
	s := &TestSuite{T: t}
	s.SetupTest()

	request := domain.CreateTaskRequest{
		Title:       "Create Task API",
		Description: "create API for create a new task",
		Status:      domain.TaskStatusTodo,
	}

	t.Run("success", func(t *testing.T) {
		s.mockRepository.On("Create",
			mock.AnythingOfType("*gorm.DB"),
			mock.AnythingOfType("*domain.Task"),
		).Return(nil).Once()

		err := s.usecase.Create(s.ctx, request)
		assert.NoError(t, err)

		s.mockRepository.AssertExpectations(t)
	})

	t.Run("create error", func(t *testing.T) {
		s.mockRepository.On("Create",
			mock.AnythingOfType("*gorm.DB"),
			mock.AnythingOfType("*domain.Task"),
		).Return(gorm.ErrInvalidData).Once()

		err := s.usecase.Create(s.ctx, request)
		assert.Error(t, err)

		s.mockRepository.AssertExpectations(t)
	})
}

func TestUpdateTask(t *testing.T) {
	s := &TestSuite{T: t}
	s.SetupTest()

	request := domain.UpdateTaskRequest{
		ID: 1,
		CreateTaskRequest: domain.CreateTaskRequest{
			Title:       "Create Task API",
			Description: "create API for create a new task",
			Status:      domain.TaskStatusTodo,
		},
	}

	t.Run("success", func(t *testing.T) {
		s.mockRepository.On("Update",
			mock.AnythingOfType("*gorm.DB"),
			mock.AnythingOfType("*domain.Task"),
		).Return(nil).Once()

		err := s.usecase.Update(s.ctx, request)
		assert.NoError(t, err)

		s.mockRepository.AssertExpectations(t)
	})

	t.Run("update error", func(t *testing.T) {
		s.mockRepository.On("Update",
			mock.AnythingOfType("*gorm.DB"),
			mock.AnythingOfType("*domain.Task"),
		).Return(gorm.ErrInvalidData).Once()

		err := s.usecase.Update(s.ctx, request)
		assert.Error(t, err)

		s.mockRepository.AssertExpectations(t)
	})
}

func TestUpdateStatusTask(t *testing.T) {
	s := &TestSuite{T: t}
	s.SetupTest()

	request := domain.UpdateStatusTaskRequest{
		ID:     1,
		Status: domain.TaskStatusDone,
	}

	t.Run("success", func(t *testing.T) {
		s.mockRepository.On("UpdateFieldByID",
			mock.AnythingOfType("*gorm.DB"),
			request.ID,
			"status",
			request.Status,
			&domain.Task{},
		).Return(nil).Once()

		err := s.usecase.UpdateStatus(s.ctx, request)
		assert.NoError(t, err)

		s.mockRepository.AssertExpectations(t)
	})

	t.Run("update error", func(t *testing.T) {
		s.mockRepository.On("UpdateFieldByID",
			mock.AnythingOfType("*gorm.DB"),
			request.ID,
			"status",
			request.Status,
			&domain.Task{},
		).Return(gorm.ErrInvalidValue).Once()

		err := s.usecase.UpdateStatus(s.ctx, request)
		assert.Error(t, err)

		s.mockRepository.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	s := &TestSuite{T: t}
	s.SetupTest()
	id := uint(1)

	t.Run("success", func(t *testing.T) {
		task := domain.Task{
			ID:          id,
			Title:       "Create Task API",
			Description: "create API for create a new task",
			Status:      domain.TaskStatusTodo,
		}

		s.mockRepository.On("FindOneByID",
			mock.AnythingOfType("*gorm.DB"),
			id,
		).Return(task, nil).Once()

		s.mockRepository.On("Delete",
			mock.AnythingOfType("*gorm.DB"),
			task,
		).Return(nil).Once()

		err := s.usecase.Delete(s.ctx, id)
		assert.NoError(t, err)

		s.mockRepository.AssertExpectations(t)
	})

	t.Run("find one by id error", func(t *testing.T) {
		s.mockRepository.On("FindOneByID",
			mock.AnythingOfType("*gorm.DB"),
			id,
		).Return(domain.Task{}, gorm.ErrRecordNotFound).Once()

		err := s.usecase.Delete(s.ctx, id)
		assert.Error(t, err)

		s.mockRepository.AssertExpectations(t)
	})

	t.Run("delete error", func(t *testing.T) {
		task := domain.Task{
			ID:          id,
			Title:       "Create Task API",
			Description: "create API for create a new task",
			Status:      domain.TaskStatusTodo,
		}

		s.mockRepository.On("FindOneByID",
			mock.AnythingOfType("*gorm.DB"),
			id,
		).Return(task, nil).Once()

		s.mockRepository.On("Delete",
			mock.AnythingOfType("*gorm.DB"),
			task,
		).Return(gorm.ErrInvalidData).Once()

		err := s.usecase.Delete(s.ctx, id)
		assert.Error(t, err)

		s.mockRepository.AssertExpectations(t)
	})
}

func TestGetOneByIDTask(t *testing.T) {
	s := &TestSuite{T: t}
	s.SetupTest()
	id := uint(1)

	t.Run("success", func(t *testing.T) {
		task := domain.Task{
			ID:          id,
			Title:       "Create Task API",
			Description: "create API for create a new task",
			Status:      domain.TaskStatusTodo,
		}

		s.mockRepository.On("FindOneByID",
			mock.AnythingOfType("*gorm.DB"),
			id,
		).Return(task, nil).Once()

		result, err := s.usecase.GetOneByID(s.ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, task, result)

		s.mockRepository.AssertExpectations(t)
	})

	t.Run("find one by id error", func(t *testing.T) {
		s.mockRepository.On("FindOneByID",
			mock.AnythingOfType("*gorm.DB"),
			id,
		).Return(domain.Task{}, gorm.ErrInvalidField).Once()

		_, err := s.usecase.GetOneByID(s.ctx, id)
		assert.Error(t, err)

		s.mockRepository.AssertExpectations(t)
	})

	t.Run("find one by id error record not found", func(t *testing.T) {
		s.mockRepository.On("FindOneByID",
			mock.AnythingOfType("*gorm.DB"),
			id,
		).Return(domain.Task{}, gorm.ErrRecordNotFound).Once()

		_, err := s.usecase.GetOneByID(s.ctx, id)
		assert.Error(t, err)

		s.mockRepository.AssertExpectations(t)
	})
}

func TestGetAllTask(t *testing.T) {
	s := &TestSuite{T: t}
	s.SetupTest()

	t.Run("success", func(t *testing.T) {
		request := domain.GetAllTaskRequest{}

		tasks := []domain.Task{
			{
				ID:          1,
				Title:       "Create Task API",
				Description: "create API for create a new task",
				Status:      domain.TaskStatusTodo,
			},
			{
				ID:          2,
				Title:       "Create Task API",
				Description: "create API for create a new task",
				Status:      domain.TaskStatusTodo,
			},
		}

		s.mockRepository.On("FindAll",
			mock.AnythingOfType("*gorm.DB"),
			request,
		).Return(tasks, nil).Once()

		result, err := s.usecase.GetAll(s.ctx, request)
		assert.NoError(t, err)
		assert.Equal(t, tasks, result)

		s.mockRepository.AssertExpectations(t)
	})

	t.Run("find all error", func(t *testing.T) {
		query := domain.GetAllTaskRequest{
			ID:          1,
			Title:       "Create Task API",
			Description: "create API for create a new task",
			Status:      domain.TaskStatusTodo,
		}

		s.mockRepository.On("FindAll",
			mock.AnythingOfType("*gorm.DB"),
			query,
		).Return([]domain.Task{}, gorm.ErrInvalidField).Once()

		_, err := s.usecase.GetAll(s.ctx, query)
		assert.Error(t, err)

		s.mockRepository.AssertExpectations(t)
	})
}
