package dto

type Mesa struct {
	Id     int64
	Numero string
	Activo bool
}

type MesaCrear struct {
	Numero string `json:"numero" binding:"required,max=20,min=1"`
}

type MesaEliminar struct {
	Id      int64 `json:"id" binding:"required,numeric"`
	Disable bool  `json:"disable" binding:"omitempty,eq=true"`
	Delete  bool  `json:"delete" binding:"omitempty,eq=true,necsfield=InnerStructField.Disable"`
}
