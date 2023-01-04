package handlers

import (
	"cosmscan-go/api/utils"
	"cosmscan-go/db"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBlockByHeight(ctx *gin.Context) {
	model := utils.MustGetDB(ctx)

	// Get the height from the URL
	height, err := strconv.ParseUint(ctx.Param("height"), 10, 64)
	if err != nil {
		sendBadRequest(ctx, errors.New("invalid number"))
		return
	}

	// Get the block from the database
	block, err := model.Block(ctx, db.BlockHeight(height))
	if err != nil {
		sendInternalErr(ctx, err)
		return
	}

	// Return the block
	ctx.JSON(http.StatusOK, block)
}
