package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-task/app/database"
	"gorm.io/gorm"
)

type (
	// Skipper defines a function to skip middleware
	Skipper func(*gin.Context) bool
)

// TransactionDatabase transaction database
func TransactionDatabase(skipper Skipper) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := &gorm.DB{}
		skip := skipper(c)

		defer func() {
			if r := recover(); r != nil {
				_ = db.Rollback()
			}
		}()

		if !skip {
			db = database.Begin()
			c.Set("db", db)
			c.Next()

			if c.Writer.Status() == http.StatusOK {
				_ = db.Commit()
			} else {
				_ = db.Rollback()
			}
		} else {
			c.Next()
		}
	}
}
