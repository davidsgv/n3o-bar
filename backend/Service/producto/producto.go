package service

import (
	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/producto"
)

type ProductoService interface {
	Save(dto.ProductoCrear)
	FindAll() []dto.Producto
	//Update(dto.CategoriaActualizar)
	FindById(id int64) *dto.Producto
}

type productoService struct{}

func NewProductoService() ProductoService {
	return &productoService{}
}

func (service productoService) Save(dto dto.ProductoCrear) {
	ProductoEntity := entity.Producto{
		Nombre:   dto.Nombre,
		Cantidad: dto.Cantidad,
		Categoria: entity.Categoria{
			Id: dto.IdCategoria,
		},
		Proveedor: entity.Tercero{
			Id: dto.IdTercero,
		},
		Historia: entity.HistoriaProducto{
			PrecioCompra: dto.PrecioCompra,
			PrecioVenta:  dto.PrecioVenta,
		},
	}

	modelo := producto.NewProductoModel()

	modelo.Save(ProductoEntity)
}

func (service productoService) FindAll() []dto.Producto {
	modelo := producto.NewProductoModel()
	ListaProductosEntity := modelo.FindAll()

	var categoriasDTO = []dto.Producto{}
	for i := 0; i < len(ListaProductosEntity); i++ {
		CategoriaAgregar := dto.Producto{
			Id:           ListaProductosEntity[i].Id,
			Nombre:       ListaProductosEntity[i].Nombre,
			Cantidad:     ListaProductosEntity[i].Cantidad,
			Fecha:        ListaProductosEntity[i].Fecha,
			Activo:       ListaProductosEntity[i].Activo,
			IdCategoria:  ListaProductosEntity[i].Categoria.Id,
			IdTercero:    ListaProductosEntity[i].Proveedor.Id,
			PrecioCompra: ListaProductosEntity[i].Historia.PrecioCompra,
			PrecioVenta:  ListaProductosEntity[i].Historia.PrecioVenta,
		}

		categoriasDTO = append(categoriasDTO, CategoriaAgregar)
	}

	return categoriasDTO
}

func (service productoService) FindById(id int64) *dto.Producto {
	modelo := producto.NewProductoModel()
	CategoriaEntity := modelo.FindById(id)

	if CategoriaEntity.Id > 0 {
		dtoResultado := dto.Producto{
			Id:           CategoriaEntity.Id,
			Nombre:       CategoriaEntity.Nombre,
			Cantidad:     CategoriaEntity.Cantidad,
			Fecha:        CategoriaEntity.Fecha,
			Activo:       CategoriaEntity.Activo,
			IdCategoria:  CategoriaEntity.Categoria.Id,
			IdTercero:    CategoriaEntity.Proveedor.Id,
			PrecioCompra: CategoriaEntity.Historia.PrecioCompra,
			PrecioVenta:  CategoriaEntity.Historia.PrecioVenta,
		}
		return &dtoResultado
	}

	return nil
}
