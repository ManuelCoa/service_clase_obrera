package interfaz

import (
	"claseobrera/domains/object/viviendas"
)

type ViviendaInterface func() ([]viviendas.Vivienda, error)

func ViviendaGetService() ViviendaInterface {
	return func() ([]viviendas.Vivienda, error) {
		return viviendas.GetViviendaService()
	}
}

func ViviendaGetServiceID(id int) ViviendaInterface {
	return func() ([]viviendas.Vivienda, error) {
		return viviendas.GetViviendaServiceID(id)
	}
}

func ViviendaPostService(data interface{}) error {
	return viviendas.PostViviendaService(data)
}

func ViviendaPutService(data interface{}) error {
	return viviendas.PutViviendaService(data)
}

func ViviendaDeleteService(id int) error {
	return viviendas.DeleteViviendaService(id)
}