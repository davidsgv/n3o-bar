package entity

type Rol struct {
	Id     int64
	Nombre string
}

type Usuario struct {
	Id       int64
	Correo   string
	Password string
	Tercero  Tercero
	Rol      Rol
}
