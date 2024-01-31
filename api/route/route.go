package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-task/api/middleware"
	"github.com/kritmet/go-task/app"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Setup setup route
func Setup(application *app.Application) {
	router := gin.Default()
	router.Use(
		middleware.TransactionDatabase(func(c *gin.Context) bool {
			return c.Request.Method == "GET"
		}),
		middleware.ErrorHandler(),
	)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	task := router.Group("/tasks")
	taskRoute(task, application)

	err := router.Run(fmt.Sprintf(":%d", application.Config.App.Port))
	if err != nil {
		panic(err)
	}
}
