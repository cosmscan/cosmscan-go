package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendErrResponse(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, gin.H{"error": err.Error()})
}

func sendBadRequest(ctx *gin.Context, err error) {
	sendErrResponse(ctx, http.StatusBadRequest, err)
}

func sendInternalErr(ctx *gin.Context, err error) {
	sendErrResponse(ctx, http.StatusInternalServerError, err)
}
