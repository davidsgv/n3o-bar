package pedido

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type FormaPagoModel interface {
	FindAll() []entity.FormaPago
}

//type terceroModel entity.Tercero
type formaPagoModel struct {
	id     sql.NullInt64
	nombre sql.NullString
}

func NewFormaPagoModel() FormaPagoModel {
	return &formaPagoModel{}
}

func (model formaPagoModel) FindAll() []entity.FormaPago {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `
		SELECT *
		FROM fpa_FormaPago
	`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.FormaPago
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.nombre)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}

		data = append(data, model.trasformToEntity())
	}

	return data
}

func (model formaPagoModel) trasformToEntity() entity.FormaPago {
	return entity.FormaPago{
		Id:     model.id.Int64,
		Nombre: model.nombre.String,
	}
}
