package entity

type Usuario struct {
	Id       int64
	Correo   string
	Password string
	Tercero  Tercero
	Rol      Rol
}
