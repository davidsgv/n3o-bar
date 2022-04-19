package dto

type Credentials struct {
	Username string `json:"username" binding:"required,max=100,min=1,email"`
	Password string `json:"password" binding:"required,max=300,min=1"`
}

//interno para generar el token en el servicio
type UsuarioInicioSesion struct {
	UsuarioId int64
	Correo    string
	Rol       string
}
