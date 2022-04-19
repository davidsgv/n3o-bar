package dto

type Caja struct {
	Id     int64
	Numero string
	Activo bool
}

type CajaCrear struct {
	Numero string `json:"numero" binding:"required,max=20,min=1"`
}

type CajaEliminar struct {
	Id      int64 `json:"id" binding:"required,numeric"`
	Disable bool  `json:"disable" binding:"omitempty,eq=true"`
	Delete  bool  `json:"delete" binding:"omitempty,eq=true,necsfield=InnerStructField.Disable"`
}
