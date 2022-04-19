package dto

import "time"

type Producto struct {
	Id           int64
	Nombre       string
	Fecha        time.Time
	Cantidad     int
	Activo       bool
	IdCategoria  int64
	IdTercero    int64
	PrecioCompra float64
	PrecioVenta  float64
}

type ProductoCrear struct {
	Nombre       string  `json:"nombre" binding:"required,max=100,min=5"`
	Cantidad     int     `json:"cantidad" binding:"required,min=1,numeric"`
	IdCategoria  int64   `json:"idCategoria" binding:"required,numeric"`
	IdTercero    int64   `json:"idTercero" binding:"required,numeric"`
	PrecioCompra float64 `json:"precioCompra" binding:"required,numeric,min=1"`
	PrecioVenta  float64 `json:"precioVenta" binding:"required,numeric,min=1"`
}

type ProductoActualizar struct {
	Id           int64   `json:"id" binding:"required,numeric"`
	Nombre       string  `json:"nombre" binding:"required,max=100,min=5"`
	Cantidad     int     `json:"cantidad" binding:"required,min=1,numeric"`
	Activo       bool    `json:"activo" binding:"required,max=1,min=0,numeric"`
	IdCategoria  int64   `json:"idCategoria" binding:"required,numeric"`
	IdTercero    int64   `json:"idTercero" binding:"required,numeric"`
	PrecioCompra float64 `json:"precioCompra" binding:"required,numeric,min=1"`
	PrecioVenta  float64 `json:"precioVenta" binding:"required,numeric,min=1"`
}
