package entity

type TipoIdentificacion struct {
	Id     int64
	Nombre string
	Codigo string
}

type Tercero struct {
	Id                 int64
	Identificacion     string
	Nombre             string
	Celular            string
	Telefono           string
	Direccion          string
	TipoIdentificacion TipoIdentificacion
}
