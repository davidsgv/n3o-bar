package service

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/model/pedido"
)

type FormaPagoService interface {
	FindAll() []dto.FormaPago
}

type formaPagoService struct{}

func NewFormaPagoService() FormaPagoService {
	return &formaPagoService{}
}

func (servicio formaPagoService) FindAll() []dto.FormaPago {
	formasPagoEntity := pedido.NewFormaPagoModel().FindAll()

	listaFormaPago := []dto.FormaPago{}
	for i := 0; i < len(formasPagoEntity); i++ {
		formaPago := dto.FormaPago{
			Id:     formasPagoEntity[i].Id,
			Nombre: formasPagoEntity[i].Nombre,
		}

		listaFormaPago = append(listaFormaPago, formaPago)
	}

	return listaFormaPago
}
