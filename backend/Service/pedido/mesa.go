package service

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/pedido"
)

type MesaService interface {
	Save(dto.MesaCrear)
	FindAll() []dto.Mesa
	FindById(id int64) *dto.Mesa
	Disable(dto.MesaEliminar)
	Delete(dto.MesaEliminar)
	//Update(dto.CategoriaActualizar)
}

type mesaService struct{}

func NewMesaService() MesaService {
	return &mesaService{}
}

func (servicio mesaService) FindAll() []dto.Mesa {
	EntidadListaMesas := pedido.NewMesasModel().FindAll()

	ListaMesas := []dto.Mesa{}
	for i := 0; i < len(EntidadListaMesas); i++ {
		Mesa := dto.Mesa{
			Id:     EntidadListaMesas[i].Id,
			Numero: EntidadListaMesas[i].Numero,
			Activo: EntidadListaMesas[i].Activo,
		}

		ListaMesas = append(ListaMesas, Mesa)
	}

	return ListaMesas
}

func (service mesaService) FindById(id int64) *dto.Mesa {
	MesaEntity := pedido.NewMesasModel().FindById(id)

	if MesaEntity.Id > 0 {
		dtoResultado := dto.Mesa{
			Id:     MesaEntity.Id,
			Numero: MesaEntity.Numero,
			Activo: MesaEntity.Activo,
		}
		return &dtoResultado
	}

	return nil
}

func (service mesaService) Save(dto dto.MesaCrear) {
	mesaEntity := entity.Mesa{
		Numero: dto.Numero,
	}

	pedido.NewMesasModel().Save(mesaEntity)
}

func (service mesaService) Disable(dto dto.MesaEliminar) {
	id := dto.Id
	pedido.NewMesasModel().Disable(id)
}

func (service mesaService) Delete(dto dto.MesaEliminar) {
	id := dto.Id
	pedido.NewMesasModel().Delete(id)
}
