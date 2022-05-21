package dto

type Categoria struct {
	Id     int64
	Nombre string
}

type CategoriaCrear struct {
	Nombre string `json:"nombre" binding:"required,max=20,min=5"`
}

type CategoriaActualizar struct {
	Id     int64  `json:"id" binding:"required,numeric"`
	Nombre string `json:"nombre" binding:"required,max=20,min=5"`
}

type CategoriaDelete struct {
	Id int64 `json:"id" binding:"required,numeric"`
}
