package controller

import (
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/tercero"
	"github.com/gin-gonic/gin"
)

type tipoIdentificacionController struct{}

func NewTipoIdentificacionController() GenericController {
	return &tipoIdentificacionController{}
}

func (controller tipoIdentificacionController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{}

	routing := router.Group("/identificacion", middlewares.AuthorizeJWT(roles))

	routing.GET("/", controller.FindAll) //find all
}

func (controller tipoIdentificacionController) FindAll(ctx *gin.Context) {
	service := service.NewTipoIdentificacionService()
	ListaTipoIdentificacion := service.FindAll()

	//datos no encontrados
	if len(ListaTipoIdentificacion) <= 0 {
		returnEmptyResultWithMessage(ctx, "No se encontraron registros")
		return
	}

	returnGoodRequest(ctx, ListaTipoIdentificacion)
}
