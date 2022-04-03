package controller

import "github.com/gin-gonic/gin"

type GenericController interface {
	Router(router *gin.RouterGroup)
}

func returnBadBinding(ctx *gin.Context, err error) {
	ctx.JSON(400, gin.H{
		"data": nil,
		"status": map[string]interface{}{
			"code":    400,
			"message": err.Error(),
		},
	})
}

func returnEmptyResultWithMessage(ctx *gin.Context, mensaje string) {
	ctx.JSON(200, gin.H{
		"data": nil,
		"status": map[string]interface{}{
			"code":    200,
			"message": mensaje,
		},
	})
}

func returnGoodRequest(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"data": data,
		"status": map[string]interface{}{
			"code":    200,
			"message": "",
		},
	})
}
