package controller

import (
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/pedido"
	"github.com/gin-gonic/gin"
)

type formaPagoController struct{}

func NewFormaPagoController() GenericController {
	return &formaPagoController{}
}

func (controller formaPagoController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{}

	routing := router.Group("/pedido/formas-pago", middlewares.AuthorizeJWT(roles))

	routing.GET("/", controller.FindAll) //find all
}

func (controller formaPagoController) FindAll(ctx *gin.Context) {
	ListaFormaPago := service.NewFormaPagoService().FindAll()

	//datos no encontrados
	if len(ListaFormaPago) <= 0 {
		returnEmptyResultWithMessage(ctx, "No se encontraron registros")
		return
	}

	returnGoodRequest(ctx, ListaFormaPago)
}
