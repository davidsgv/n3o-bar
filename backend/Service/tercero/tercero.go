package service

import (
	dto "github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/entity"
	terceroModel "github.com/davidsgv/n3o-bar/model/tercero"
)

//interfaz para uso externo
type TerceroService interface {
	Save(dto.TerceroCrear)
	FindAll() []dto.Tercero
	FindById(id int64) *dto.Tercero
	Update(dto.TerceroActualizar)
}

type terceroService struct{}

func NewTerceroService() TerceroService {
	return &terceroService{}
}

func (service terceroService) Update(dto dto.TerceroActualizar) {
	terceroEntity := entity.Tercero{
		Nombre:         dto.Nombre,
		Celular:        dto.Celular,
		Telefono:       dto.Telefono,
		Direccion:      dto.Direccion,
		Identificacion: dto.Identificacion,
		TipoIdentificacion: entity.TipoIdentificacion{
			Id: dto.TipoIndetificacionId,
		},
	}

	modelo := terceroModel.NewTerceroModel()

	modelo.Update(terceroEntity)
}

func (service terceroService) Save(dto dto.TerceroCrear) {
	terceroEntity := entity.Tercero{
		Nombre:         dto.Nombre,
		Celular:        dto.Celular,
		Telefono:       dto.Telefono,
		Direccion:      dto.Direccion,
		Identificacion: dto.Identificacion,
		TipoIdentificacion: entity.TipoIdentificacion{
			Id: dto.TipoIndetificacionId,
		},
	}

	modelo := terceroModel.NewTerceroModel()

	modelo.Save(terceroEntity)
}

func (service terceroService) FindAll() []dto.Tercero {
	modelo := terceroModel.NewTerceroModel()
	ListaterceroEntity := modelo.FindAll()

	var listaTercerosDTO = []dto.Tercero{}
	for i := 0; i < len(ListaterceroEntity); i++ {
		terceroAgregar := dto.Tercero{
			Id:                       ListaterceroEntity[i].Id,
			Identificacion:           ListaterceroEntity[i].Identificacion,
			Nombre:                   ListaterceroEntity[i].Nombre,
			Celular:                  ListaterceroEntity[i].Celular,
			Telefono:                 ListaterceroEntity[i].Telefono,
			Direccion:                ListaterceroEntity[i].Direccion,
			TipoIndetificacionId:     ListaterceroEntity[i].TipoIdentificacion.Id,
			TipoIdentificacionNombre: ListaterceroEntity[i].TipoIdentificacion.Nombre,
			TipoIdentificacionCodigo: ListaterceroEntity[i].TipoIdentificacion.Codigo,
		}

		listaTercerosDTO = append(listaTercerosDTO, terceroAgregar)
	}

	return listaTercerosDTO
}

func (service terceroService) FindById(id int64) *dto.Tercero {
	modelo := terceroModel.NewTerceroModel()

	TerceroEntity := modelo.FindById(id)

	if TerceroEntity.Id > 0 {
		dtoResultado := dto.Tercero{
			Id:                       TerceroEntity.Id,
			Identificacion:           TerceroEntity.Identificacion,
			Nombre:                   TerceroEntity.Nombre,
			Celular:                  TerceroEntity.Celular,
			Telefono:                 TerceroEntity.Telefono,
			Direccion:                TerceroEntity.Direccion,
			TipoIndetificacionId:     TerceroEntity.TipoIdentificacion.Id,
			TipoIdentificacionNombre: TerceroEntity.TipoIdentificacion.Nombre,
			TipoIdentificacionCodigo: TerceroEntity.TipoIdentificacion.Codigo,
		}
		return &dtoResultado
	}

	return nil
}
