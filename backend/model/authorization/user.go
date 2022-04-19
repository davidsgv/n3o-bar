package authorization

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type UserModel interface {
	Create(entity.Usuario)
	FindById(int) entity.Usuario
	FindAll() []entity.Usuario
}

//type terceroModel entity.Tercero
type userModel struct {
	usuId     sql.NullInt64
	Correo    sql.NullString
	RolId     sql.NullInt64
	TerId     sql.NullInt64
	terNombre sql.NullString
}

func NewUserModel() UserModel {
	return &userModel{}
}

func (model userModel) Create(usuario entity.Usuario) {
	statement := "CALL spCrearUsuario" +
		"($1, $2, $3, $4, $5, $6"

	if usuario.Tercero.Celular != "" {
		statement += ",$7"
		usuario.Tercero.Celular = "null"
	}
	if usuario.Tercero.Telefono != "" {
		statement += ",$8"
		usuario.Tercero.Telefono = "null"
	}
	if usuario.Tercero.Direccion != "" {
		statement += ",$9"
		usuario.Tercero.Direccion = "null"
	}
	statement += ")"

	sql := database.Init()

	stmt, err := sql.Prepare(statement)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(usuario.Tercero.Identificacion, usuario.Tercero.TipoIdentificacion.Id,
		usuario.Tercero.Nombre, usuario.Correo, usuario.Password, usuario.Rol.Id,
		usuario.Tercero.Celular, usuario.Tercero.Telefono, usuario.Tercero.Direccion)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model userModel) FindAll() []entity.Usuario {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `SELECT usu_Id, usu_Correo, rol_Id, ter.ter_Id, ter_Nombre
				FROM usu_Usuario usu
				INNER JOIN ter_Tercero ter
					ON ter.ter_Id = usu.ter_Id`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Usuario
	for rows.Next() {
		scanErr := rows.Scan(&model.usuId, &model.Correo, &model.RolId,
			&model.TerId, &model.terNombre)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}

		data = append(data, model.trasformToEntity())
	}

	return data
}

func (model userModel) FindById(id int) entity.Usuario {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `SELECT usu_Id, usu_Correo, rol_Id, ter.ter_Id, ter_Nombre
				FROM usu_Usuario usu
				INNER JOIN ter_Tercero ter
					ON ter.ter_Id = usu.ter_Id
				WHERE usu_Id = $1`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query(id)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data entity.Usuario
	for rows.Next() {
		scanErr := rows.Scan(&model.usuId, &model.Correo, &model.RolId,
			&model.TerId, &model.terNombre)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}
		data = model.trasformToEntity()
	}

	return data
}

func (model userModel) trasformToEntity() entity.Usuario {
	return entity.Usuario{
		Id:     model.usuId.Int64,
		Correo: model.Correo.String,
		Tercero: entity.Tercero{
			Id:     model.TerId.Int64,
			Nombre: model.terNombre.String,
		},
		Rol: entity.Rol{
			Id: model.RolId.Int64,
		},
	}
}
