package dto

type UsuarioCrear struct {
	Tercero TerceroCrear `json:"tercero" binding:"omitempty"`
	Credentials
	RolId int64 `json:"idRol" binding:"required,min=1,max=3,numeric"`
}

type UsuarioBuscarById struct {
	Id int64 `uri:"id" binding:"numeric"`
}

type UsuarioDesactivar struct {
	Id     int64 `json:"id" binding:"required,numeric"`
	Activo bool  `json:"activo"`
}

type Usuario struct {
	Id        int64
	Correo    string
	Activo    bool
	RolNombre string
	TerNombre string
}
