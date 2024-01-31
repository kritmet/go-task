package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-task/domain"
)

// NewTaskController is a function for create task controller
func NewTaskController(usecase domain.TaskUsecase) domain.TaskController {
	return &taskController{
		usecase: usecase,
	}
}

type taskController struct {
	usecase domain.TaskUsecase
}

// Create Create
// @Tags Task
// @Summary Create
// @Description create task
// @Accept json
// @Produce json
// @Param request body domain.CreateTaskRequest true "input create request"
// @Success 200 {object} domain.Message
// @Failure 400 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /tasks [post]
func (c *taskController) Create(ctx *gin.Context) {
	request := domain.CreateTaskRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = c.usecase.Create(request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domain.NewSuccessMessage())
}

// Update Update
// @Tags Task
// @Summary Update
// @Description update task
// @Accept json
// @Produce json
// @Param id path int true "task id"
// @Param request body domain.UpdateTaskRequest true "input update request"
// @Success 200 {object} domain.Message
// @Failure 400 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /tasks/{id} [put]
func (c *taskController) Update(ctx *gin.Context) {
	var request domain.UpdateTaskRequest
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	request.ID = uint(i)

	err = ctx.Bind(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = c.usecase.Update(request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domain.NewSuccessMessage())
}

// UpdateStatus Update status
// @Tags Task
// @Summary UpdateStatus
// @Description update task status
// @Accept json
// @Produce json
// @Param id path int true "task id"
// @Param request body domain.UpdateStatusTaskRequest true "input update request"
// @Success 200 {object} domain.Message
// @Failure 400 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /tasks/{id}/status [patch]
func (c *taskController) UpdateStatus(ctx *gin.Context) {
	var request domain.UpdateStatusTaskRequest
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	request.ID = uint(i)

	err = ctx.Bind(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = c.usecase.UpdateStatus(request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domain.NewSuccessMessage())
}

// Delete delete
// @Tags Task
// @Summary Delete
// @Description delete task
// @Accept json
// @Produce json
// @Param id path int true "task id"
// @Success 200 {object} domain.Message
// @Failure 400 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /tasks/{id} [delete]
func (c *taskController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = c.usecase.Delete(uint(i))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domain.NewSuccessMessage())
}

// GetOneByID get one by id
// @Tags Task
// @Summary GetOneByID
// @Description get task by id
// @Accept json
// @Produce json
// @Param id path int true "task id"
// @Success 200 {object} domain.Task
// @Failure 404 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /tasks/{id} [get]
func (c *taskController) GetOneByID(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	task, err := c.usecase.GetOneByID(uint(i))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// GetAll get all
// @Tags Task
// @Summary GetAll
// @Description get all task
// @Accept json
// @Produce json
// @Param request query domain.GetAllTaskRequest true "query for get all"
// @Success 200 {array} domain.Task
// @Failure 400 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /tasks [get]
func (c *taskController) GetAll(ctx *gin.Context) {
	request := domain.GetAllTaskRequest{}
	err := ctx.Bind(request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	tasks, err := c.usecase.GetAll(request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
