package controller

import (
	"errors"
	"strconv"

	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/middlewares"
	service "github.com/davidsgv/n3o-bar/service/producto"
	"github.com/gin-gonic/gin"
)

type productoController struct{}

func NewProductoController() GenericController {
	return &productoController{}
}

func (controller productoController) Router(router *gin.RouterGroup) {
	//roles que pueden acceder a esta ruta
	roles := []string{"Administrador"}

	routing := router.Group("/producto", middlewares.AuthorizeJWT(roles))

	routing.GET("/", controller.FindAll)     //find all
	routing.GET("/:id", controller.FindById) //find by id
	routing.POST("/", controller.Save)
	// routing.PUT("/", controller.Update)
}

func (controller productoController) FindAll(ctx *gin.Context) {
	listaProductos := service.NewProductoService().FindAll()

	if len(listaProductos) <= 0 {
		returnEmptyResultWithMessage(ctx, "No hay registros")
		return
	}

	returnGoodRequest(ctx, listaProductos)
}

func (controller productoController) FindById(ctx *gin.Context) {
	service := service.NewProductoService()

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

func (controller productoController) Save(ctx *gin.Context) {
	service := service.NewProductoService()
	dto := dto.ProductoCrear{}
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	service.Save(dto)

	returnGoodRequest(ctx, nil)
}

// func (controller categoriaController) Update(ctx *gin.Context) {
// 	service := service.NewCategoriaService()
// 	dto := dto.CategoriaActualizar{}
// 	err := ctx.ShouldBindJSON(&dto)

// 	if err != nil {
// 		returnBadBinding(ctx, err)
// 		return
// 	}

// 	service.Update(dto)

// 	returnEmptyResultWithMessage(ctx, "Registro actualizado correctamente")
// }
