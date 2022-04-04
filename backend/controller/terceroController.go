package controller

import (
	"errors"
	"strconv"

	service "github.com/davidsgv/n3o-bar/Service/tercero"
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/middlewares"
	"github.com/gin-gonic/gin"
)

type terceroController struct{}

func NewTerceroController() GenericController {
	return &terceroController{}
}

func (controller terceroController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{"Administrador"}

	routing := router.Group("/tercero", middlewares.AuthorizeJWT(roles))

	routing.GET("/", controller.FindAll)     //find all
	routing.GET("/:id", controller.FindById) //find by id
	routing.POST("/", controller.Save)
	routing.PUT("/", controller.Update)
}

func (controller terceroController) Update(ctx *gin.Context) {
	service := service.NewTerceroService()
	dto := dto.TerceroActualizar{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	service.Update(dto)

	returnEmptyResultWithMessage(ctx, "Registro actualizado correctamente")
}

func (controller terceroController) FindAll(ctx *gin.Context) {
	listaTercero := service.NewTerceroService().FindAll()

	if len(listaTercero) <= 0 {
		returnEmptyResultWithMessage(ctx, "No hay registros")
		return
	}

	returnGoodRequest(ctx, listaTercero)
}

func (controller terceroController) FindById(ctx *gin.Context) {
	service := service.NewTerceroService()

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		returnBadBinding(ctx, errors.New("El id debe ser un numero"))
		return
	}

	dtoRespuesta := service.FindById(id)
	if dtoRespuesta == nil {
		returnEmptyResultWithMessage(ctx, "El id no existe")
		return
	}

	returnGoodRequest(ctx, dtoRespuesta)
}

func (controller terceroController) Save(ctx *gin.Context) {
	service := service.NewTerceroService()
	dto := dto.TerceroCrear{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	service.Save(dto)

	returnGoodRequest(ctx, nil)
}
