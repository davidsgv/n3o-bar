package producto

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type ProductoModel interface {
	Save(entity.Producto)
	FindAll() []entity.Producto
	FindById(int64) entity.Producto
	//Update(entity.Producto)
}

//type terceroModel entity.Tercero
type productoModel struct {
	id       sql.NullInt64
	nombre   sql.NullString
	fecha    sql.NullTime
	cantidad sql.NullInt32
	activo   sql.NullBool

	idTercero sql.NullInt64

	precioCompra sql.NullFloat64
	precioVenta  sql.NullFloat64

	idCategoria     sql.NullInt64
	nombreCategoria sql.NullString
}

func NewProductoModel() ProductoModel {
	return &productoModel{}
}

func (model productoModel) Save(entity entity.Producto) {
	sql := database.Init()

	stmt, err := sql.Prepare(`CALL spCrearProducto(
								$1
								, $2
								, $3
								, $4
								, $5
								, $6
							)`)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(entity.Nombre, entity.Cantidad, entity.Categoria.Id,
		entity.Proveedor.Id, entity.Historia.PrecioCompra, entity.Historia.PrecioVenta)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model productoModel) FindAll() []entity.Producto {
	//inicia la conexión
	sql := database.Init()

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(`
		SELECT 
			pro.pro_Id
			, pro_Nombre
			, pro_Cantidad
			, pro_Activo
			, ter_Id
			, cat_Nombre
			, hip_PrecioCompra::numeric
			, hip_PrecioVenta::numeric
			, hip_FechaCreacion
		FROM pro_Producto pro
		INNER JOIN cat_Categoria cat
			ON cat.cat_Id = pro.cat_id 
		INNER JOIN hip_historiaproducto hip
			ON pro.pro_id = hip.pro_id
	`)
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Producto
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.nombre, &model.cantidad, &model.activo, &model.idTercero,
			&model.nombreCategoria, &model.precioCompra, &model.precioVenta, &model.fecha)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		} else {
			data = append(data, model.trasformToEntity())
		}
	}

	return data
}

func (model productoModel) FindById(id int64) entity.Producto {
	//inicia la conexión
	database := database.Init()

	//realiza la consulta a la base de datos
	row, err := database.Query(`
		SELECT 
			pro.pro_Id
			, pro_Nombre
			, pro_Cantidad
			, pro_Activo
			, ter_Id
			, cat_Nombre
			, hip_PrecioCompra::numeric
			, hip_PrecioVenta::numeric
			, hip_FechaCreacion
		FROM pro_Producto pro
		INNER JOIN cat_Categoria cat
			ON cat.cat_Id = pro.cat_id 
		INNER JOIN hip_historiaproducto hip
			ON pro.pro_id = hip.pro_id
		WHERE pro.pro_Id = $1
	`, id)
	defer row.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en las variables
	for row.Next() {
		scanErr := row.Scan(&model.id, &model.nombre, &model.cantidad, &model.activo, &model.idTercero,
			&model.nombreCategoria, &model.precioCompra, &model.precioVenta, &model.fecha)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}
	}

	//trasforma la consulta en la entidad
	data := model.trasformToEntity()

	return data
}

// func (model productoModel) Update(entity entity.Producto) {
// 	sql := database.Init()

// 	_, err := sql.Exec(`UPDATE cat_Categoria
// 						SET cat_Nombre = $1
// 						WHERE cat_Id = $2`,
// 		entity.Nombre, entity.Id)
// 	defer sql.Close()

// 	if err != nil {
// 		panic(err)
// 	}
// }

func (model productoModel) trasformToEntity() entity.Producto {
	// precioCompra, err := strconv.ParseFloat(model.precioCompra.String, 64)
	// if err != nil {
	// 	panic(errors.New("Error al convertir money to decimal"))
	// }
	// precioVenta, err := strconv.ParseFloat(model.precioVenta.String, 64)
	// if err != nil {
	// 	panic(errors.New("Error al convertir money to decimal"))
	// }

	return entity.Producto{
		Id:       model.id.Int64,
		Nombre:   model.nombre.String,
		Fecha:    model.fecha.Time,
		Cantidad: int(model.cantidad.Int32),
		Activo:   model.activo.Bool,
		Categoria: entity.Categoria{
			Id:     model.idCategoria.Int64,
			Nombre: model.nombreCategoria.String,
		},
		Proveedor: entity.Tercero{
			Id: model.idTercero.Int64,
		},
		Historia: entity.HistoriaProducto{
			PrecioCompra: model.precioCompra.Float64,
			PrecioVenta:  model.precioVenta.Float64,
		},
	}
}
