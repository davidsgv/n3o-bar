package producto

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type CategoriaModel interface {
	Save(entity.Categoria)
	FindAll() []entity.Categoria
	FindById(int64) entity.Categoria
	Update(entity.Categoria)
	Delete(int64) error
}

//type terceroModel entity.Tercero
type categoriaModel struct {
	id     sql.NullInt64
	nombre sql.NullString
}

func NewCategoriaModel() CategoriaModel {
	return &categoriaModel{}
}

func (model categoriaModel) Save(entity entity.Categoria) {
	sql := database.Init()

	stmt, err := sql.Prepare("INSERT INTO cat_Categoria" +
		"(cat_nombre) VALUES ($1)")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(entity.Nombre)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model categoriaModel) FindAll() []entity.Categoria {
	//inicia la conexión
	sql := database.Init()

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare("SELECT * FROM cat_Categoria")
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Categoria
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.nombre)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		} else {
			data = append(data, model.trasformToEntity())
		}

	}

	return data
}

func (model categoriaModel) FindById(id int64) entity.Categoria {
	//inicia la conexión
	database := database.Init()

	//realiza la consulta a la base de datos
	row, err := database.Query("SELECT * FROM cat_Categoria WHERE cat_Id = $1", id)
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

func (model categoriaModel) Update(entity entity.Categoria) {
	sql := database.Init()

	_, err := sql.Exec(`UPDATE cat_Categoria
						SET cat_Nombre = $1
						WHERE cat_Id = $2`,
		entity.Nombre, entity.Id)
	defer sql.Close()

	if err != nil {
		panic(err)
	}
}

func (model categoriaModel) Delete(id int64) error {
	sql := database.Init()

	_, err := sql.Exec(`DELETE
						FROM cat_categoria
						WHERE cat_categoria.cat_id = $1`,
		id)
	defer sql.Close()

	if strings.Contains(err.Error(), "update o delete en «cat_categoria» viola la llave foránea «fk_pro_cat» en la tabla «pro_producto") {
		return errors.New("Existe un producto con esta categoria, no es posible eliminarla")
	}

	if err != nil {
		fmt.Printf(err.Error())
		panic(err)
	}

	return nil
}

func (model categoriaModel) trasformToEntity() entity.Categoria {
	return entity.Categoria{
		Id:     model.id.Int64,
		Nombre: model.nombre.String,
	}
}
