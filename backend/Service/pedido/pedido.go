package service

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/pedido"
)

type PedidoService interface {
	Save(dto.PedidoCrear)
	// FindAll() []dto.Mesa
	// FindById(id int64) *dto.Mesa
	// Disable(dto.MesaEliminar)
	// Delete(dto.MesaEliminar)
	//Update(dto.CategoriaActualizar)
}

type pedidoService struct{}

func NewPedidoService() PedidoService {
	return &pedidoService{}
}

func (service pedidoService) Save(dto dto.PedidoCrear) {
	//valida que la mesa este activa
	mesa := NewMesaService().FindById(dto.IdMesa)
	if mesa.Activo == false {
		panic("La mesa no esta activa")
	}

	entity := entity.Pedido{
		Mesa: entity.Mesa{
			Id: dto.IdMesa,
		},
		Mesero: entity.Usuario{
			Id: dto.Mesero,
		},
	}
	//crea el pedido
	pedido.NewPedidoModel().Create(entity)
}
