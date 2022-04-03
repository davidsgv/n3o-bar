package main

import (
	"io"
	"os"

	service "github.com/davidsgv/n3o-bar/Service"
	"github.com/davidsgv/n3o-bar/controller"
	"github.com/davidsgv/n3o-bar/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	loginService service.AuthorizationService = service.NewAuthorizationService()
	//jwtService   service.JWTService   = service.NewJWTService()

	//loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

var controllers []controller.GenericController = []controller.GenericController{
	controller.NewTipoIdentificacionController(), //tipo identificacion
	controller.NewTerceroController(),            //tercero
	controller.NewAuthorizationController(),      //authorizacion
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(),
		middlewares.Logger())
	//middlewares.Logger(), middlewares.BasicAuth())

	// server.POST("/login", func(ctx *gin.Context) {
	// 	token := loginController.Login(ctx)
	// 	fmt.Println("TOKEN: " + token)

	// 	if token != "" {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": token,
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, nil)
	// 	}
	// })

	//Se registran las rutas del api
	apiRoutes := server.Group("/api/")
	for i := 0; i < len(controllers); i++ {
		controllers[i].Router(apiRoutes)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run(":" + port)
}
