package handlers

import (
	"cosmscan-go/api/utils"
	"github.com/gin-gonic/gin"
)

func GetAllChains(c *gin.Context) {
	model := utils.MustGetDB(c)
	chains, err := model.AllChains(c)
	if err != nil {
		sendInternalErr(c, err)
		return
	}

	c.JSON(200, chains)
}
