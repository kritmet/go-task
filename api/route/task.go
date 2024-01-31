package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-task/api/controller"
	"github.com/kritmet/go-task/app"
	"github.com/kritmet/go-task/repository"
	"github.com/kritmet/go-task/usecase"
)

func taskRoute(route *gin.RouterGroup, application *app.Application) {
	uc := usecase.NewTaskUsecase(
		application,
		repository.NewTaskRepository(),
	)
	ct := controller.NewTaskController(uc)

	route.POST("", ct.Create)
	route.PUT("/:id", ct.Update)
	route.PATCH("/:id/status", ct.UpdateStatus)
	route.DELETE("/:id", ct.Delete)
	route.GET("/:id", ct.GetOneByID)
	route.GET("", ct.GetAll)
}
