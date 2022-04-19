package entity

import "time"

type HistoriaProducto struct {
	Id           int64
	Fecha        time.Time
	PrecioCompra float64
	PrecioVenta  float64
}
