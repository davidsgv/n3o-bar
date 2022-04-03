package service

import (
	"github.com/davidsgv/n3o-bar/dto"
	terceroModel "github.com/davidsgv/n3o-bar/model/tercero"
)

type TipoIdentificacionService interface {
	FindAll() []dto.TipoIdentificacion
}

type tipoIdentificacionService struct {
	*dto.TipoIdentificacion
}

func NewTipoIdentificacionService() TipoIdentificacionService {
	return &tipoIdentificacionService{}
}

func (service *tipoIdentificacionService) FindAll() []dto.TipoIdentificacion {
	modelo := terceroModel.NewTipoIdentificacionModel()

	ListaTipoIdentificacionEntity := modelo.FindAll()

	var listaTipoIdentificacionDTO = []dto.TipoIdentificacion{}
	for i := 0; i < len(ListaTipoIdentificacionEntity); i++ {
		TipoIdentificacion := dto.TipoIdentificacion{
			Id:     ListaTipoIdentificacionEntity[i].Id,
			Codigo: ListaTipoIdentificacionEntity[i].Codigo,
			Nombre: ListaTipoIdentificacionEntity[i].Nombre,
		}

		listaTipoIdentificacionDTO = append(listaTipoIdentificacionDTO, TipoIdentificacion)
	}

	return listaTipoIdentificacionDTO
}
