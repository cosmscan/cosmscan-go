package utils

import (
	"cosmscan-go/api/constants"
	"cosmscan-go/db"

	"github.com/gin-gonic/gin"
)

func MustGetDB(c *gin.Context) db.DB {
	database, ok := c.Get(constants.CONTEXT_DB)
	if !ok {
		panic("failed to get database instance from context, key \"" + constants.CONTEXT_DB + "\" not found")
	}
	return database.(db.DB)
}
