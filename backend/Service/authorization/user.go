package authorization

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/authorization"
)

//interfaz para uso externo
type UserService interface {
	CreateUser(dto.UsuarioCrear)
	FindById(int) *dto.Usuario
	FindAll() []dto.Usuario
}

//implementa la interfaz para devolver este objeto
type userService struct{}

//contructor
func NewUserService() UserService {
	return &userService{}
}

func (service userService) CreateUser(dtoEntrada dto.UsuarioCrear) {
	model := authorization.NewUserModel()
	entity := entity.Usuario{
		Tercero: entity.Tercero{
			Identificacion: dtoEntrada.Tercero.Identificacion,
			Nombre:         dtoEntrada.Tercero.Nombre,
			Celular:        dtoEntrada.Tercero.Celular,
			Telefono:       dtoEntrada.Tercero.Telefono,
			Direccion:      dtoEntrada.Tercero.Direccion,
			TipoIdentificacion: entity.TipoIdentificacion{
				Id: dtoEntrada.Tercero.TipoIndetificacionId,
			},
		},
		Correo:   dtoEntrada.Username,
		Password: dtoEntrada.Password,
		Rol: entity.Rol{
			Id: dtoEntrada.RolId,
		},
	}

	model.Create(entity)
}
func (service userService) FindById(id int) *dto.Usuario {
	modelo := authorization.NewUserModel()

	userEntity := modelo.FindById(id)

	if userEntity.Id > 0 {
		dtoResultado := dto.Usuario{
			Id:        userEntity.Id,
			Correo:    userEntity.Correo,
			TerNombre: userEntity.Tercero.Nombre,
			TerId:     userEntity.Tercero.Id,
			RolId:     userEntity.Rol.Id,
		}
		return &dtoResultado
	}

	return nil
}
func (service userService) FindAll() []dto.Usuario {
	modelo := authorization.NewUserModel()
	ListaEntity := modelo.FindAll()

	var listaDTO = []dto.Usuario{}
	for i := 0; i < len(ListaEntity); i++ {
		usuario := dto.Usuario{
			Id:        ListaEntity[i].Id,
			Correo:    ListaEntity[i].Correo,
			TerNombre: ListaEntity[i].Tercero.Nombre,
			TerId:     ListaEntity[i].Tercero.Id,
			RolId:     ListaEntity[i].Rol.Id,
		}

		listaDTO = append(listaDTO, usuario)
	}

	return listaDTO
}
