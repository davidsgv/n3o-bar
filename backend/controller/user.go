package controller

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/middlewares"
	"github.com/davidsgv/n3o-bar/service/authorization"
	"github.com/gin-gonic/gin"
)

type userController struct{}

func NewUserController() GenericController {
	return &userController{}
}

func (controlador userController) Router(router *gin.RouterGroup) {
	var roles []string = []string{"Administrador"}

	userRouter := router.Group("/user", middlewares.AuthorizeJWT(roles))
	userRouter.POST("", controlador.create)

}

func (controller *userController) create(ctx *gin.Context) {
	dto := dto.UsuarioCrear{}
	err := ctx.ShouldBind(&dto)

	//si falla el binding
	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	//crear usuario
	authorizationService := authorization.NewUserService()
	authorizationService.CreateUser(dto)

	//devuelve el token
	returnGoodRequest(ctx, "usuario creado")
}
