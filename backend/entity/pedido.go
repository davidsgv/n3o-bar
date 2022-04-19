package entity

import "time"

type Pedido struct {
	Id           int64
	FormaPago    FormaPago
	EstadoPedido EstadoPedido
	Mesa         Mesa
	Caja         Caja
	Fecha        time.Time
	Mesero       Usuario
	Cajero       Usuario
	Cliente      Tercero
}
