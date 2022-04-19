package service

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/model/pedido"
)

type EstadoPedidoService interface {
	FindAll() []dto.EstadoPedido
	FindById(id int64) *dto.EstadoPedido
}

type estadoPedidoService struct{}

func NewEstadoPedidoService() EstadoPedidoService {
	return &estadoPedidoService{}
}

func (servicio estadoPedidoService) FindAll() []dto.EstadoPedido {
	EntidadListaEstados := pedido.NewEstadoPedidoModel().FindAll()

	ListaEstados := []dto.EstadoPedido{}
	for i := 0; i < len(EntidadListaEstados); i++ {
		Estado := dto.EstadoPedido{
			Id:     EntidadListaEstados[i].Id,
			Nombre: EntidadListaEstados[i].Nombre,
		}

		ListaEstados = append(ListaEstados, Estado)
	}

	return ListaEstados
}

//No esta probado
func (service estadoPedidoService) FindById(id int64) *dto.EstadoPedido {
	EstadoEntity := pedido.NewEstadoPedidoModel().FindById(id)

	if EstadoEntity.Id > 0 {
		dtoResultado := dto.EstadoPedido{
			Id:     EstadoEntity.Id,
			Nombre: EstadoEntity.Nombre,
		}
		return &dtoResultado
	}

	return nil
}
