package interfaz

import (
	"claseobrera/domains/object/partido_politico"
)

type PartidoPoliticoInterface func() ([]partido_politico.PartidoPolitico, error)

func PartidoPoliticoGetService() PartidoPoliticoInterface {
	return func() ([]partido_politico.PartidoPolitico, error) {
		return partido_politico.GetPartidoPoliticoService()
	}
}

func PartidoPoliticoGetServiceID(id int) PartidoPoliticoInterface {
	return func() ([]partido_politico.PartidoPolitico, error) {
		return partido_politico.GetPartidoPoliticoServiceID(id)
	}
}

func PartidoPoliticoPostService(data interface{}) error {
	return partido_politico.PostPartidoPoliticoService(data)
}

func PartidoPoliticoPutService(data interface{}) error {
	return partido_politico.PutPartidoPoliticoService(data)
}

func PartidoPoliticoDeleteService(id int) error {
	return partido_politico.DeletePartidoPoliticoService(id)
}