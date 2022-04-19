package service

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/pedido"
)

type CajaService interface {
	Save(dto.CajaCrear)
	FindAll() []dto.Caja
	FindById(id int64) *dto.Caja
	Disable(dto.CajaEliminar)
	Delete(dto.CajaEliminar)
	//Update(dto.CategoriaActualizar)
}

type cajaService struct{}

func NewCajaService() CajaService {
	return &cajaService{}
}

func (servicio cajaService) FindAll() []dto.Caja {
	EntidadListaCajas := pedido.NewCajaModel().FindAll()

	ListaCajas := []dto.Caja{}
	for i := 0; i < len(EntidadListaCajas); i++ {
		caja := dto.Caja{
			Id:     EntidadListaCajas[i].Id,
			Numero: EntidadListaCajas[i].Numero,
			Activo: EntidadListaCajas[i].Activo,
		}

		ListaCajas = append(ListaCajas, caja)
	}

	return ListaCajas
}

func (service cajaService) FindById(id int64) *dto.Caja {
	CajaEntity := pedido.NewCajaModel().FindById(id)

	if CajaEntity.Id > 0 {
		dtoResultado := dto.Caja{
			Id:     CajaEntity.Id,
			Numero: CajaEntity.Numero,
			Activo: CajaEntity.Activo,
		}
		return &dtoResultado
	}

	return nil
}

func (service cajaService) Save(dto dto.CajaCrear) {
	cajaEntity := entity.Caja{
		Numero: dto.Numero,
	}

	pedido.NewCajaModel().Save(cajaEntity)
}

func (service cajaService) Disable(dto dto.CajaEliminar) {
	id := dto.Id
	pedido.NewCajaModel().Disable(id)
}

func (service cajaService) Delete(dto dto.CajaEliminar) {
	id := dto.Id
	pedido.NewCajaModel().Delete(id)
}
