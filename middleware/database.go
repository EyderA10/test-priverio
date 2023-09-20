package middleware

import (
	"technical-test/priverion/utils"

	"github.com/gin-gonic/gin"
)

func DatabaseMiddleware(db *utils.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
