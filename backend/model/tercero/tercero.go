package terceroModel

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type TerceroModel interface {
	Save(entity.Tercero)
	FindAll() []entity.Tercero
	FindById(int64) entity.Tercero
	Update(entity.Tercero)
}

//type terceroModel entity.Tercero
type terceroModel struct {
	id                   sql.NullInt64
	idTipoIdentificacion sql.NullInt64
	identificacion       sql.NullString
	nombre               sql.NullString
	celular              sql.NullString
	telefono             sql.NullString
	direccion            sql.NullString
}

func NewTerceroModel() TerceroModel {
	return &terceroModel{}
}

func (model terceroModel) Save(entity entity.Tercero) {
	sql := database.Init()

	stmt, err := sql.Prepare("INSERT INTO ter_tercero" +
		"(ter_nombre, ter_celular, ter_direccion, ter_identificacion, ter_telefono, tid_id) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(entity.Nombre, entity.Celular, entity.Direccion, entity.Identificacion, entity.Telefono, entity.TipoIdentificacion.Id)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model terceroModel) FindAll() []entity.Tercero {
	//inicia la conexión
	sql := database.Init()

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare("SELECT * FROM ter_Tercero")
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Tercero
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.identificacion, &model.idTipoIdentificacion,
			&model.nombre, &model.celular, &model.telefono, &model.direccion)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		} else {
			data = append(data, model.trasformToEntity())
		}

	}

	return data
}

func (model terceroModel) FindById(id int64) entity.Tercero {
	//inicia la conexión
	database := database.Init()

	//realiza la consulta a la base de datos
	row, err := database.Query("SELECT * FROM ter_Tercero WHERE ter_Id = $1", id)
	defer row.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en las variables
	for row.Next() {
		scanErr := row.Scan(&model.id, &model.identificacion, &model.idTipoIdentificacion,
			&model.nombre, &model.celular, &model.telefono, &model.direccion)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}
	}

	//trasforma la consulta en la entidad
	data := model.trasformToEntity()

	return data
}

func (model terceroModel) Update(entity entity.Tercero) {
	sql := database.Init()

	_, err := sql.Exec("UPDATE ter_tercero"+
		"SET ter_nombre = '%s'"+
		", ter_celular = '%s'"+
		", ter_direccion = '%s'"+
		", ter_identificacion = '%s'"+
		", ter_telefono = '%s'"+
		", tid_id = '%d'"+
		" FROM ter_Tercero",
		entity.Nombre, entity.Celular, entity.Direccion, entity.Identificacion, entity.Telefono, entity.TipoIdentificacion.Id)
	defer sql.Close()

	if err != nil {
		panic(err)
	}
}

func (model terceroModel) trasformToEntity() entity.Tercero {
	return entity.Tercero{
		Id:             model.id.Int64,
		Identificacion: model.identificacion.String,
		Nombre:         model.nombre.String,
		Celular:        model.celular.String,
		Telefono:       model.telefono.String,
		Direccion:      model.direccion.String,
		TipoIdentificacion: entity.TipoIdentificacion{
			Id: model.idTipoIdentificacion.Int64,
		},
	}
}
