package dto

type UsuarioCrear struct {
	Tercero TerceroCrear `json:"tercero" binding:"omitempty"`
	Credentials
	RolId int64 `json:"idRol" binding:"required,min=1,max=3,numeric"`
}

type UsuarioBuscarById struct {
	Id int64 `uri:"id" binding:"numeric"`
}

type Usuario struct {
	Id        int64
	Correo    string
	RolId     int64
	TerId     int64
	TerNombre string
}
