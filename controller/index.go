package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context){
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "index",
	})
	return
}
