package main

import (
	"io"
	"os"

	"github.com/davidsgv/n3o-bar/controller"
	"github.com/davidsgv/n3o-bar/middlewares"
	"github.com/gin-gonic/gin"
)

//services
// var (
// 	loginService authorization.AuthorizationService = authorization.NewAuthorizationService()
// )

//controllers
var controllers []controller.GenericController = []controller.GenericController{
	//Administracion
	controller.NewAuthorizationController(), //authorizacion
	controller.NewRolController(),           //Roles

	//Tercero
	controller.NewTipoIdentificacionController(), //tipo identificacion
	controller.NewTerceroController(),            //tercero

	//Productos
	controller.NewCategoriaController(), //categorias
	controller.NewProductoController(),  //productos

	//Pedidos
	controller.NewMesaController(),         //Mesas
	controller.NewCajaController(),         //Cajas
	controller.NewEstadoPedidoController(), //Estados Pedidos
	controller.NewFormaPagoController(),    //Formas de pago
	controller.NewPedidoController(),       //Pedidos
}

func main() {
	//Logger setup
	setupLogOutput()

	server := gin.New()

	//Logger and Recovery for error
	server.Use(gin.Recovery(),
		middlewares.Logger())

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

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
