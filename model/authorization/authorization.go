package authorization

import (
	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type AuthorizationModel interface {
	Create(usuario entity.Usuario) error
	Login(usuario entity.Usuario) *entity.Usuario
}

//type terceroModel entity.Tercero
type authorizationModel struct{}

func NewAuthorizationModel() AuthorizationModel {
	return &authorizationModel{}
}

func (model *authorizationModel) Create(usuario entity.Usuario) error {
	sql := database.Init()

	stmt, err := sql.Prepare("CALL spCrearUsuario" +
		"($1, $2, $3, $4, $5, $6) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(usuario.Tercero.Identificacion, usuario.Tercero.TipoIdentificacion.Id, usuario.Tercero.Nombre, usuario.Correo, usuario.Password, usuario.Rol.Id)
	defer sql.Close()
	if err != nil {
		return err
	}
	return nil
}

func (model *authorizationModel) Login(usuario entity.Usuario) *entity.Usuario {
	//inicia la conexión
	sql := database.Init()

	//realiza la consulta a la base de datos
	row, err := sql.Query("SELECT * FROM fnIniciarSesion($1,$2)", usuario.Correo, usuario.Password)
	defer row.Close()
	if err != nil {
		panic(err)
	}

	var data entity.Usuario
	//obtiene los datos y los guarda en la variable data
	for row.Next() {
		scanErr := row.Scan(&data.Id, &data.Correo, &data.Rol.Id, &data.Rol.Nombre)
		if scanErr != nil {
			panic(scanErr)
		}
	}

	return &data
}
