package interfaz

import (
	"claseobrera/domains/object/estados"
)

type EstadosInterface func() ([]estados.Estado, error)

func EstadoService() EstadosInterface {
	return func() ([]estados.Estado, error) {
		return estados.GetEstadosService()
	}
}

func EstadoIDService(id int) EstadosInterface {
	return func() ([]estados.Estado, error) {
		return estados.GetEstadoIDService(id)
	}
}

func EstadoPostService(data interface{}) error {
	return estados.PostEstadoService(data)
}

func EstadoPutService(data interface{}) error {
	return estados.PutEstadoService(data)
}

func EstadoDeleteService(id int) error {
	return estados.DeleteEstadoService(id)
}
