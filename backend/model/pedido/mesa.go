package pedido

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type MesaModel interface {
	FindAll() []entity.Mesa
	FindById(int64) entity.Mesa
	Save(entity.Mesa)
	Disable(int64)
	Delete(int64)
	Update(entity.Mesa)
}

//type terceroModel entity.Tercero
type mesaModel struct {
	id     sql.NullInt64
	nombre sql.NullString
	activo sql.NullBool
}

func NewMesasModel() MesaModel {
	return &mesaModel{}
}

func (model mesaModel) FindAll() []entity.Mesa {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `
				SELECT mes_Id, mes_Numero, mes_Activo
				FROM mes_Mesa
				WHERE mes_Eliminado = false
				`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Mesa
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.nombre, &model.activo)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}

		data = append(data, model.trasformToEntity())
	}

	return data
}

func (model mesaModel) FindById(id int64) entity.Mesa {
	//inicia la conexión
	database := database.Init()

	//realiza la consulta a la base de datos
	row, err := database.Query(`
		SELECT mes_Id, mes_Numero, mes_Activo
		FROM mes_Mesa pro
		WHERE mes_Id = $1
	`, id)
	defer row.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en las variables
	for row.Next() {
		scanErr := row.Scan(&model.id, &model.nombre, &model.activo)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}
	}

	//trasforma la consulta en la entidad
	data := model.trasformToEntity()

	return data
}

func (model mesaModel) Save(entity entity.Mesa) {
	sql := database.Init()

	stmt, err := sql.Prepare(`
		INSERT INTO mes_Mesa (mes_Numero)
		VALUES ($1)
	`)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(entity.Numero)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model mesaModel) Disable(id int64) {
	sql := database.Init()

	stmt, err := sql.Prepare(`
		UPDATE mes_Mesa
		SET mes_Activo = false
		WHERE mes_Id = $1
	`)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model mesaModel) Delete(id int64) {
	sql := database.Init()

	stmt, err := sql.Prepare(`
		UPDATE mes_Mesa
		SET mes_Eliminado = true,
			mes_Activo = false
		WHERE mes_Id = $1
	`)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model mesaModel) Update(entity entity.Mesa) {
	sql := database.Init()

	stmt, err := sql.Prepare(`
		UPDATE mes_Mesa
		SET mes_Numero = $1,
			mes_Activo = $2
		WHERE mes_Id = $3
	`)

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(entity.Numero, entity.Activo, entity.Id)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model mesaModel) trasformToEntity() entity.Mesa {
	return entity.Mesa{
		Id:     model.id.Int64,
		Numero: model.nombre.String,
		Activo: model.activo.Bool,
	}
}
