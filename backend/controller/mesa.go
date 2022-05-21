package controller

import (
	"errors"
	"strconv"

	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/pedido"
	"github.com/gin-gonic/gin"
)

type mesaController struct{}

func NewMesaController() GenericController {
	return &mesaController{}
}

func (controller mesaController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{"Administrador", "Mesero"}

	routing := router.Group("/mesa", middlewares.AuthorizeJWT(roles))

	routing.GET("/", controller.FindAll)     //find all
	routing.GET("/:id", controller.FindById) //find by id
	routing.POST("/", controller.Save)       //create
	routing.DELETE("/", controller.Delete)   //delete
	routing.PUT("/", controller.Update)      //update
	// routing.PUT("/", controller.Update)
}

func (controller mesaController) FindAll(ctx *gin.Context) {
	listaMesas := service.NewMesaService().FindAll()

	if len(listaMesas) <= 0 {
		returnEmptyResultWithMessage(ctx, "No hay registros")
		return
	}

	returnGoodRequest(ctx, listaMesas)
}

func (controller mesaController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		returnBadBinding(ctx, errors.New("El id debe ser un numero"))
		return
	}

	dtoRespuesta := service.NewMesaService().FindById(id)
	if dtoRespuesta == nil {
		returnEmptyResultWithMessage(ctx, "El id no existe")
		return
	}

	returnGoodRequest(ctx, dtoRespuesta)
}

func (controller mesaController) Save(ctx *gin.Context) {
	dto := dto.MesaCrear{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	service.NewMesaService().Save(dto)
	returnGoodRequest(ctx, nil)
}

func (controller mesaController) Delete(ctx *gin.Context) {
	dto := dto.MesaEliminar{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	if dto.Disable {
		service.NewMesaService().Disable(dto)
		returnGoodRequest(ctx, nil)
		return
	}
	if dto.Delete {
		service.NewMesaService().Delete(dto)
		returnGoodRequest(ctx, nil)
		return
	}

	returnBadBinding(ctx, errors.New("Formato incorrecto"))
}

func (controller mesaController) Update(ctx *gin.Context) {
	dto := dto.MesaActualizar{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	service.NewMesaService().Update(dto)
	returnGoodRequest(ctx, nil)
}
