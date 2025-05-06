package interfaz

import (
	"claseobrera/domains/object/tipo_suelo_vivienda"
)

type TipoSueloViviendaInterface func() ([]tipo_suelo_vivienda.TipoSueloVivienda, error)

func TipoSueloViviendaGetService() TipoSueloViviendaInterface {
	return func() ([]tipo_suelo_vivienda.TipoSueloVivienda, error) {
		return tipo_suelo_vivienda.GetTiposSueloService()
	}
}

func TipoSueloViviendaGetServiceID(id int) TipoSueloViviendaInterface {
	return func() ([]tipo_suelo_vivienda.TipoSueloVivienda, error) {
		return tipo_suelo_vivienda.GetTiposSueloServiceID(id)
	}
}

func TipoSueloViviendaPostService(data interface{}) error {
	return tipo_suelo_vivienda.PostTipoSueloService(data)
}

func TipoSueloViviendaPutService(data interface{}) error {
	return tipo_suelo_vivienda.PutTipoSueloService(data)
}

func TipoSueloViviendaDeleteService(id int) error {
	return tipo_suelo_vivienda.DeleteTipoSueloService(id)
}