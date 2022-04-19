package authorization

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/model/authorization"
)

//interfaz para uso externo
type RolService interface {
	FindAll() []dto.Rol
}

//implementa la interfaz para devolver este objeto
type rolService struct{}

//contructor
func NewRolService() RolService {
	return &rolService{}
}

func (service rolService) FindAll() []dto.Rol {
	modelo := authorization.NewRolModel()
	ListaEntity := modelo.FindAll()

	var listaDTO = []dto.Rol{}
	for i := 0; i < len(ListaEntity); i++ {
		usuario := dto.Rol{
			Id:     ListaEntity[i].Id,
			Nombre: ListaEntity[i].Nombre,
		}

		listaDTO = append(listaDTO, usuario)
	}

	return listaDTO
}
