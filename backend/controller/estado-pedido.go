package controller

import (
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/pedido"
	"github.com/gin-gonic/gin"
)

type estadoPedidoController struct{}

func NewEstadoPedidoController() GenericController {
	return &estadoPedidoController{}
}

func (controller estadoPedidoController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{}

	routing := router.Group("/pedido/estados", middlewares.AuthorizeJWT(roles))

	routing.GET("/", controller.FindAll) //find all
}

func (controller estadoPedidoController) FindAll(ctx *gin.Context) {
	ListaTipoIdentificacion := service.NewEstadoPedidoService().FindAll()

	//datos no encontrados
	if len(ListaTipoIdentificacion) <= 0 {
		returnEmptyResultWithMessage(ctx, "No se encontraron registros")
		return
	}

	returnGoodRequest(ctx, ListaTipoIdentificacion)
}
