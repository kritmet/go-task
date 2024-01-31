package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-task/app/config"
)

// ErrorHandler error handler middleware
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx.Next()

		lastErr := ctx.Errors.Last()
		for _, err := range ctx.Errors {
			if err == lastErr {
				if e, ok := err.Err.(config.Result); ok {
					ctx.AbortWithStatusJSON(
						e.Code,
						e,
					)
					return
				}

				ctx.AbortWithStatusJSON(
					http.StatusInternalServerError,
					config.NewMessageError(ctx.Writer.Status(), err.Error()),
				)
			}
		}
	}
}
