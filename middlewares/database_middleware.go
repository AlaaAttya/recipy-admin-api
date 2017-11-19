package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
)

// Database  Add database connection as middleware
func Database(connection *bongo.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DBConnection", connection)
		c.Next()
	}
}
