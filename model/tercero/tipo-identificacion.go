package terceroModel

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type TipoIdentificacionModel interface {
	FindAll() []entity.TipoIdentificacion
}

//type terceroModel entity.Tercero
type tipoIdentificacionModel struct {
	id     sql.NullInt64
	nombre sql.NullString
	codigo sql.NullString
}

func NewTipoIdentificacionModel() TipoIdentificacionModel {
	return &tipoIdentificacionModel{}
}

func (model *tipoIdentificacionModel) FindAll() []entity.TipoIdentificacion {
	//inicia la conexión
	sql := database.Init()

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare("SELECT * FROM tid_TipoIdentificacion")
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.TipoIdentificacion
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.nombre, &model.codigo)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		} else {
			data = append(data, model.trasformToEntity())
		}

	}

	return data
}

func (model tipoIdentificacionModel) trasformToEntity() entity.TipoIdentificacion {
	return entity.TipoIdentificacion{
		Id:     model.id.Int64,
		Nombre: model.nombre.String,
		Codigo: model.codigo.String,
	}
}
