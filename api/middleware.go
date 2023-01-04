package api

import (
	"cosmscan-go/api/constants"
	"cosmscan-go/db"

	"github.com/gin-gonic/gin"
)

// MiddlewareDatabaseContext adds the database instance to the gin.Context
// Every handler must have a database instance
func MiddlewareDatabaseContext(database db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constants.CONTEXT_DB, database)
		c.Next()
	}
}
