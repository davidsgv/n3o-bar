package controller

import (
	"errors"
	"strconv"

	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/pedido"
	"github.com/gin-gonic/gin"
)

type cajaController struct{}

func NewCajaController() GenericController {
	return &cajaController{}
}

func (controller cajaController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{"Administrador", "Cajero"}

	routing := router.Group("/caja", middlewares.AuthorizeJWT(roles))

	routing.GET("/", controller.FindAll)     //find all
	routing.GET("/:id", controller.FindById) //find by id
	routing.POST("/", controller.Save)       //create
	routing.DELETE("/", controller.Delete)   //delete
	// routing.PUT("/", controller.Update)
}

func (controller cajaController) FindAll(ctx *gin.Context) {
	listaMesas := service.NewCajaService().FindAll()

	if len(listaMesas) <= 0 {
		returnEmptyResultWithMessage(ctx, "No hay registros")
		return
	}

	returnGoodRequest(ctx, listaMesas)
}

func (controller cajaController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		returnBadBinding(ctx, errors.New("El id debe ser un numero"))
		return
	}

	dtoRespuesta := service.NewCajaService().FindById(id)
	if dtoRespuesta == nil {
		returnEmptyResultWithMessage(ctx, "El id no existe")
		return
	}

	returnGoodRequest(ctx, dtoRespuesta)
}

func (controller cajaController) Save(ctx *gin.Context) {
	dto := dto.CajaCrear{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	service.NewCajaService().Save(dto)
	returnGoodRequest(ctx, nil)
}

func (controller cajaController) Delete(ctx *gin.Context) {
	dto := dto.CajaEliminar{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	if dto.Disable {
		service.NewCajaService().Disable(dto)
		returnGoodRequest(ctx, nil)
		return
	}
	if dto.Delete {
		service.NewCajaService().Delete(dto)
		returnGoodRequest(ctx, nil)
		return
	}

	returnBadBinding(ctx, errors.New("Formato incorrecto"))
}
