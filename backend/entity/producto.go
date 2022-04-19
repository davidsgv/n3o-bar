package entity

import "time"

type Producto struct {
	Id        int64
	Nombre    string
	Fecha     time.Time
	Cantidad  int
	Activo    bool
	Categoria Categoria
	Proveedor Tercero
	Historia  HistoriaProducto
}
