package dto

type PedidoCrear struct {
	IdMesa int64 `json:"idmesa" binding:"required,numeric"`
	Mesero int64
}
