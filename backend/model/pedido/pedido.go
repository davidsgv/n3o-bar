package pedido

import (
	"database/sql"
	"errors"

	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/database"
)

type PedidoModel interface {
	FindAll() []entity.Pedido
	Create(entity.Pedido)
	FindByEstado(int64) []entity.Pedido
	// FindById(int64) entity.Pedido
	// Save(entity.Pedido)
	// Update(entity.Pedido)
}

//type terceroModel entity.Tercero
type pedidoModel struct {
	id        sql.NullInt64
	fecha     sql.NullTime
	mesero    sql.NullString
	cajero    sql.NullString
	cliente   sql.NullString
	mesa      sql.NullString
	estado    sql.NullString
	formaPago sql.NullString
	caja      sql.NullString
}

func NewPedidoModel() PedidoModel {
	return &pedidoModel{}
}

func (model pedidoModel) Create(entity entity.Pedido) {
	//inicia la conexión
	sql := database.Init()

	stmt, err := sql.Prepare(`
		INSERT INTO ped_pedido (mes_Id, epe_Id, ped_Mesero)
		VALUES ($1, 1, $2)
	`)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(entity.Mesa.Id, entity.Mesero.Id)
	defer sql.Close()
	if err != nil {
		panic(err)
	}
}

func (model pedidoModel) FindAll() []entity.Pedido {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `
			SELECT 
				ped_Id
				, ped_fechaCreacion AS FECHAPEDIDO
				, usuMesero.usu_Correo AS MESERO
				, usuCajero.usu_Correo AS CAJERO
				, ter.ter_Nombre AS CLIENTE
				, mes.mes_numero AS MESA
				, epe_Nombre AS ESTADO 
				, fpa.fpa_nombre AS FORMAPAGO
				, caj.caj_Numero AS CAJA
			FROM ped_pedido ped
			INNER JOIN usu_usuario usuMesero
				ON ped.ped_mesero = usuMesero.usu_id
			LEFT JOIN usu_usuario usuCajero
				ON ped.ped_cajero = usuCajero.usu_id
			LEFT JOIN ter_tercero ter
				ON ped.ped_cliente = ter.ter_id
			INNER JOIN epe_estadopedido epe
				ON ped.epe_id = epe.epe_id
			LEFT JOIN fpa_formapago fpa
				ON ped.fpa_id = fpa.fpa_id
			INNER JOIN mes_mesa mes
				ON ped.mes_id = mes.mes_id
			LEFT JOIN caj_caja caj
				ON caj.caj_id = ped.caj_id
	`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Pedido
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.fecha, &model.mesero, &model.cajero,
			&model.cliente, &model.mesa, &model.estado, &model.formaPago, &model.caja)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}

		data = append(data, model.trasformToEntity())
	}

	return data
}

func (model pedidoModel) FindByEstado(idEstado int64) []entity.Pedido {
	//inicia la conexión
	sql := database.Init()

	//consulta
	statement := `
			SELECT 
				ped_Id
				, ped_fechaCreacion AS FECHAPEDIDO
				, usuMesero.usu_Correo AS MESERO
				, usuCajero.usu_Correo AS CAJERO
				, ter.ter_Nombre AS CLIENTE
				, mes.mes_numero AS MESA
				, epe_Nombre AS ESTADO 
				, fpa.fpa_nombre AS FORMAPAGO
				, caj.caj_Numero AS CAJA
			FROM ped_pedido ped
			INNER JOIN usu_usuario usuMesero
				ON ped.ped_mesero = usuMesero.usu_id
			LEFT JOIN usu_usuario usuCajero
				ON ped.ped_cajero = usuCajero.usu_id
			LEFT JOIN ter_tercero ter
				ON ped.ped_cliente = ter.ter_id
			INNER JOIN epe_estadopedido epe
				ON ped.epe_id = epe.epe_id
			LEFT JOIN fpa_formapago fpa
				ON ped.fpa_id = fpa.fpa_id
			INNER JOIN mes_mesa mes
				ON ped.mes_id = mes.mes_id
			LEFT JOIN caj_caja caj
				ON caj.caj_id = ped.caj_id
			WHERE ped.epe_id = $1
	`

	//realiza la consulta a la base de datos
	prepareDB, err := sql.Prepare(statement)
	rows, err := prepareDB.Query(idEstado)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	//obtiene los datos y los guarda en la variable data
	var data []entity.Pedido
	for rows.Next() {
		scanErr := rows.Scan(&model.id, &model.fecha, &model.mesero, &model.cajero,
			&model.cliente, &model.mesa, &model.estado, &model.formaPago, &model.caja)
		if scanErr != nil {
			panic(errors.New("Error al obtener los datos"))
		}

		data = append(data, model.trasformToEntity())
	}

	return data
}

// func (model mesaModel) FindByEstado(id int64) entity.Mesa {
// 	//inicia la conexión
// 	database := database.Init()

// 	//realiza la consulta a la base de datos
// 	row, err := database.Query(`
// 		SELECT mes_Id, mes_Numero, mes_Activo
// 		FROM mes_Mesa pro
// 		WHERE mes_Id = $1
// 	`, id)
// 	defer row.Close()
// 	if err != nil {
// 		panic(err)
// 	}

// 	//obtiene los datos y los guarda en las variables
// 	for row.Next() {
// 		scanErr := row.Scan(&model.id, &model.nombre, &model.activo)
// 		if scanErr != nil {
// 			panic(errors.New("Error al obtener los datos"))
// 		}
// 	}

// 	//trasforma la consulta en la entidad
// 	data := model.trasformToEntity()

// 	return data
// }

// func (model mesaModel) Save(entity entity.Mesa) {
// 	sql := database.Init()

// 	stmt, err := sql.Prepare(`
// 		INSERT INTO mes_Mesa (mes_Numero)
// 		VALUES ($1)
// 	`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = stmt.Exec(entity.Numero)
// 	defer sql.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (model mesaModel) Disable(id int64) {
// 	sql := database.Init()

// 	stmt, err := sql.Prepare(`
// 		UPDATE mes_Mesa
// 		SET mes_Activo = false
// 		WHERE mes_Id = $1
// 	`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = stmt.Exec(id)
// 	defer sql.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (model mesaModel) Delete(id int64) {
// 	sql := database.Init()

// 	stmt, err := sql.Prepare(`
// 		UPDATE mes_Mesa
// 		SET mes_Eliminado = true,
// 			mes_Activo = false
// 		WHERE mes_Id = $1
// 	`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = stmt.Exec(id)
// 	defer sql.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func (model pedidoModel) trasformToEntity() entity.Pedido {
	return entity.Pedido{
		Id: model.id.Int64,
		FormaPago: entity.FormaPago{
			Nombre: model.formaPago.String,
		},
		EstadoPedido: entity.EstadoPedido{
			Nombre: model.estado.String,
		},
		Mesa: entity.Mesa{
			Numero: model.mesa.String,
		},
		Caja: entity.Caja{
			Numero: model.caja.String,
		},
		Fecha: model.fecha.Time,
		Mesero: entity.Usuario{
			Correo: model.mesero.String,
		},
		Cajero: entity.Usuario{
			Correo: model.cajero.String,
		},
		Cliente: entity.Tercero{
			Nombre: model.cliente.String,
		},
	}
}
