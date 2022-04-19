package pedido

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type CajaModel interface {
	FindAll() []entity.Caja
	FindById(int64) entity.Caja
	Save(entity.Caja)
	Disable(int64)
	Delete(int64)
}

//type terceroModel entity.Tercero
type cajaModel struct {
	id     sql.NullInt64
	nombre sql.NullString
	activo sql.NullBool
}

func NewCajaModel() CajaModel {
	return &cajaModel{}
}

func (model cajaModel) FindAll() []entity.Caja {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `
				SELECT caj_Id, caj_Numero, caj_Activo
				FROM caj_caja
				WHERE caj_Eliminado = false
				`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Caja
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.nombre, &model.activo)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}

		data = append(data, model.trasformToEntity())
	}

	return data
}

func (model cajaModel) FindById(id int64) entity.Caja {
	//inicia la conexión
	database := database.Init()

	//realiza la consulta a la base de datos
	row, err := database.Query(`
		SELECT caj_Id, caj_Numero, caj_Activo
		FROM caj_caja
		WHERE caj_Id = $1
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

func (model cajaModel) Save(entity entity.Caja) {
	sql := database.Init()

	stmt, err := sql.Prepare(`
		INSERT INTO caj_Caja (caj_Numero)
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

func (model cajaModel) Disable(id int64) {
	sql := database.Init()

	stmt, err := sql.Prepare(`
		UPDATE caj_Caja
		SET caj_Activo = false
		WHERE caj_Id = $1
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

func (model cajaModel) Delete(id int64) {
	sql := database.Init()

	stmt, err := sql.Prepare(`
		UPDATE caj_Caja
		SET caj_Eliminado = true,
			caj_Activo = false
		WHERE caj_Id = $1
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

func (model cajaModel) trasformToEntity() entity.Caja {
	return entity.Caja{
		Id:     model.id.Int64,
		Numero: model.nombre.String,
		Activo: model.activo.Bool,
	}
}
