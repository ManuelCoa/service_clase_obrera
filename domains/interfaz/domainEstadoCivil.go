package interfaz

import (
	"claseobrera/domains/object/estado_civil"
)

type EstadoCivilInterface func() ([]estado_civil.EstadoCivil, error)

func EstadoCivilGetService() EstadoCivilInterface {
	return func() ([]estado_civil.EstadoCivil, error) {
		return estado_civil.GetEstadoCivilService()
	}
}

func EstadoCivilGetServiceID(id int) EstadoCivilInterface {
	return func() ([]estado_civil.EstadoCivil, error) {
		return estado_civil.GetEstadoCivilServiceID(id)
	}
}

func EstadoCivilPostService(data interface{}) error {
	return estado_civil.PostEstadoCivilService(data)
}

func EstadoCivilPutService(data interface{}) error {
	return estado_civil.PutEstadoCivilService(data)
}

func EstadoCivilDeleteService(id int) error {
	return estado_civil.DeleteEstadoCivilService(id)
}