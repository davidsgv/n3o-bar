package dto

type Tercero struct {
	Id                       int64
	Identificacion           string
	Nombre                   string
	Celular                  string
	Telefono                 string
	Direccion                string
	TipoIndetificacionId     int64
	TipoIdentificacionNombre string
	TipoIdentificacionCodigo string
}

type TerceroCrear struct {
	Identificacion       string `json:"identificacion" binding:"required,max=20,min=10,numeric"`
	Nombre               string `json:"nombre" binding:"required,max=50"`
	Celular              string `json:"celular" binding:"omitempty,max=20,min=10,numeric"`
	Telefono             string `json:"telefono" binding:"omitempty,max=20,min=7,numeric"`
	Direccion            string `json:"direccion" binding:"omitempty,max=50,min=10"`
	TipoIndetificacionId int64  `json:"tipoIdentificacion" binding:"required,min=1,max=3"`
}

type TerceroActualizar struct {
	Id                   int64  `json:"id" binding:"required,gte=1,numeric"`
	Identificacion       string `json:"identificacion" binding:"required,max=20,min=10,numeric"`
	Nombre               string `json:"nombre" binding:"required,max=50"`
	Celular              string `json:"celular" binding:"required,max=20,min=10,numeric"`
	Telefono             string `json:"telefono" binding:"required,max=20,min=7,numeric"`
	Direccion            string `json:"direccion" binding:"required,max=50,min=10"`
	TipoIndetificacionId int64  `json:"tipoIdentificacion" binding:"required,min=1,max=3"`
}

type TerceroBuscarById struct {
	Id int64 `uri:"id" binding:"numeric"`
}
