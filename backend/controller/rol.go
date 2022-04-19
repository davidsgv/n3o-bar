package controller

import (
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/authorization"
	"github.com/gin-gonic/gin"
)

type rolController struct{}

func NewRolController() GenericController {
	return &rolController{}
}

func (controller rolController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{"Administrador"}

	routing := router.Group("/rol", middlewares.AuthorizeJWT(roles))

	routing.GET("/", controller.FindAll) //find all
}

func (controller rolController) FindAll(ctx *gin.Context) {
	service := service.NewRolService()
	listaRoles := service.FindAll()

	//datos no encontrados
	if len(listaRoles) <= 0 {
		returnEmptyResultWithMessage(ctx, "No se encontraron registros")
		return
	}

	returnGoodRequest(ctx, listaRoles)
}
