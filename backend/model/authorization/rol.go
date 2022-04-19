package authorization

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type RolModel interface {
	FindAll() []entity.Rol
}

//type terceroModel entity.Tercero
type rolModel struct {
	id     sql.NullInt64
	nombre sql.NullString
}

func NewRolModel() RolModel {
	return &rolModel{}
}

func (model rolModel) FindAll() []entity.Rol {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `SELECT * FROM rol_Rol`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Rol
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.nombre)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}

		data = append(data, model.trasformToEntity())
	}

	return data
}

func (model rolModel) trasformToEntity() entity.Rol {
	return entity.Rol{
		Id:     model.id.Int64,
		Nombre: model.nombre.String,
	}
}
