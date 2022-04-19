package controller

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/pedido"
	"github.com/gin-gonic/gin"
)

type pedidoController struct{}

func NewPedidoController() GenericController {
	return &pedidoController{}
}

func (controller pedidoController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{"Administrador", "Mesero"}

	routing := router.Group("/pedido", middlewares.AuthorizeJWT(roles))

	routing.POST("/", controller.Save) //find all
	// routing.GET("/:id", controller.FindById) //find by id
	// routing.GET("/", controller.FindAll)       //create
	// routing.DELETE("/", controller.Delete)   //delete
	// routing.PUT("/", controller.Update)
}

func (controller pedidoController) Save(ctx *gin.Context) {
	dto := dto.PedidoCrear{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	//asigna el id del usuario al dto
	dto.Mesero = int64(ctx.GetFloat64("userId"))
	service.NewPedidoService().Save(dto)
	returnGoodRequest(ctx, nil)
}
