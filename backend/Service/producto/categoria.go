package service

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/entity"
	productoModel "github.com/davidsgv/n3o-bar/model/producto"
)

type CategoriaService interface {
	Save(dto.CategoriaCrear)
	FindAll() []dto.Categoria
	Update(dto.CategoriaActualizar)
	FindById(id int64) *dto.Categoria
}

type categoriaService struct{}

func NewCategoriaService() CategoriaService {
	return &categoriaService{}
}

func (service categoriaService) Save(dto dto.CategoriaCrear) {
	categoriaEntity := entity.Categoria{
		Nombre: dto.Nombre,
	}

	modelo := productoModel.NewCategoriaModel()

	modelo.Save(categoriaEntity)
}

func (service categoriaService) Update(dto dto.CategoriaActualizar) {
	categoriaEntity := entity.Categoria{
		Id:     dto.Id,
		Nombre: dto.Nombre,
	}

	modelo := productoModel.NewCategoriaModel()
	modelo.Update(categoriaEntity)
}

func (service categoriaService) FindAll() []dto.Categoria {
	modelo := productoModel.NewCategoriaModel()
	ListaCategoriasEntity := modelo.FindAll()

	var categoriasDTO = []dto.Categoria{}
	for i := 0; i < len(ListaCategoriasEntity); i++ {
		CategoriaAgregar := dto.Categoria{
			Id:     ListaCategoriasEntity[i].Id,
			Nombre: ListaCategoriasEntity[i].Nombre,
		}

		categoriasDTO = append(categoriasDTO, CategoriaAgregar)
	}

	return categoriasDTO
}

func (service categoriaService) FindById(id int64) *dto.Categoria {
	modelo := productoModel.NewCategoriaModel()
	CategoriaEntity := modelo.FindById(id)

	if CategoriaEntity.Id > 0 {
		dtoResultado := dto.Categoria{
			Id:     CategoriaEntity.Id,
			Nombre: CategoriaEntity.Nombre,
		}
		return &dtoResultado
	}

	return nil
}
