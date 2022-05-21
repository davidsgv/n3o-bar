package controller

import (
	"errors"
	"strconv"

	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/producto"
	"github.com/gin-gonic/gin"
)

type categoriaController struct{}

func NewCategoriaController() GenericController {
	return &categoriaController{}
}

func (controller categoriaController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{"Administrador"}

	routing := router.Group("/categoria", middlewares.AuthorizeJWT(roles))

	//todos los roles
	routing.GET("/", controller.FindAll)     //find all
	routing.GET("/:id", controller.FindById) //find by id

	//administrados
	routing.POST("/", controller.Save)
	routing.PUT("/", controller.Update)
	routing.DELETE("/", controller.Delete)
}

func (controller categoriaController) FindAll(ctx *gin.Context) {
	listaCategorias := service.NewCategoriaService().FindAll()

	if len(listaCategorias) <= 0 {
		returnEmptyResultWithMessage(ctx, "No hay registros")
		return
	}

	returnGoodRequest(ctx, listaCategorias)
}

func (controller categoriaController) Update(ctx *gin.Context) {
	service := service.NewCategoriaService()
	dto := dto.CategoriaActualizar{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	service.Update(dto)

	returnEmptyResultWithMessage(ctx, "Registro actualizado correctamente")
}

func (controller categoriaController) FindById(ctx *gin.Context) {
	service := service.NewCategoriaService()

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

func (controller categoriaController) Save(ctx *gin.Context) {
	service := service.NewCategoriaService()
	dto := dto.CategoriaCrear{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	service.Save(dto)

	returnGoodRequest(ctx, nil)
}

func (controller categoriaController) Delete(ctx *gin.Context) {
	service := service.NewCategoriaService()
	dto := dto.CategoriaDelete{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	err = service.Delete(dto)

	if err != nil {
		returnEmptyResultWithMessage(ctx, err.Error())
		return
	}
	returnEmptyResultWithMessage(ctx, "Categoria eliminada correctamente")
}
