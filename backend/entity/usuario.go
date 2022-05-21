package entity

type Usuario struct {
	Id       int64
	Correo   string
	Password string
	Activo   bool
	Tercero  Tercero
	Rol      Rol
}
