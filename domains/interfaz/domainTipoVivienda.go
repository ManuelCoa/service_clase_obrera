package interfaz

import (
	"claseobrera/domains/object/tipo_vivienda"
)

type TipoViviendaInterface func() ([]tipo_vivienda.TipoVivienda, error)

func TipoViviendaGetService() TipoViviendaInterface {
	return func() ([]tipo_vivienda.TipoVivienda, error) {
		return tipo_vivienda.GetTiposViviendaService()
	}
}

func TipoViviendaGetServiceID(id int) TipoViviendaInterface {
	return func() ([]tipo_vivienda.TipoVivienda, error) {
		return tipo_vivienda.GetTiposViviendaServiceID(id)
	}
}

func TipoViviendaPostService(data interface{}) error {
	return tipo_vivienda.PostTipoViviendaService(data)
}

func TipoViviendaPutService(data interface{}) error {
	return tipo_vivienda.PutTipoViviendaService(data)
}

func TipoViviendaDeleteService(id int) error {
	return tipo_vivienda.DeleteTipoViviendaService(id)
}