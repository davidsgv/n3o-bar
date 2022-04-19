package pedido

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type EstadoPedidoModel interface {
	FindAll() []entity.EstadoPedido
	FindById(int64) entity.EstadoPedido
}

//type terceroModel entity.Tercero
type estadoPedidoModel struct {
	id     sql.NullInt64
	nombre sql.NullString
}

func NewEstadoPedidoModel() EstadoPedidoModel {
	return &estadoPedidoModel{}
}

func (model estadoPedidoModel) FindAll() []entity.EstadoPedido {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `
				SELECT *
				FROM epe_EstadoPedido
				`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.EstadoPedido
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.nombre)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}

		data = append(data, model.trasformToEntity())
	}

	return data
}

func (model estadoPedidoModel) FindById(id int64) entity.EstadoPedido {
	//inicia la conexión
	database := database.Init()

	//realiza la consulta a la base de datos
	row, err := database.Query(`
		SELECT *
		FROM epe_EstadoPedido
		WHERE epe_Id = $1
	`, id)
	defer row.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en las variables
	for row.Next() {
		scanErr := row.Scan(&model.id, &model.nombre)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}
	}

	//trasforma la consulta en la entidad
	data := model.trasformToEntity()

	return data
}

func (model estadoPedidoModel) trasformToEntity() entity.EstadoPedido {
	return entity.EstadoPedido{
		Id:     model.id.Int64,
		Nombre: model.nombre.String,
	}
}
