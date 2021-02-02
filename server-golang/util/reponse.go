package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, status int, code int, data interface{}, msg string) {
	ctx.JSON(status, gin.H{"code": code, "data": data, "message": msg})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
