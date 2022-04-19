package controller

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/service/authorization"
	"github.com/gin-gonic/gin"
)

type authorizationController struct{}

func NewAuthorizationController() GenericController {
	return &authorizationController{}
}

func (controlador authorizationController) Router(router *gin.RouterGroup) {
	router.POST("login/", controlador.login)
}

func (controller *authorizationController) login(ctx *gin.Context) {
	//hacer binding con el dto
	dto := dto.Credentials{}
	err := ctx.ShouldBind(&dto)

	//si falla el binding
	if err != nil {
		returnBadBinding(ctx, err)
		return
	}

	//iniciar sesion
	authorizationService := authorization.NewAuthorizationService()
	dtoRespuesta, resultado := authorizationService.Login(dto)

	//Si el inicio de sesión falla
	if resultado == false {
		returnEmptyResultWithMessage(ctx, "Credenciales invalidas")
		return
	}

	//generar el token
	token := authorizationService.GenerateToken(*dtoRespuesta)

	//devuelve el token
	returnGoodRequest(ctx, token)
}
