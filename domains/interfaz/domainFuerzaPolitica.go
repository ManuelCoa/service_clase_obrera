package interfaz

import (
	"claseobrera/domains/object/fuerza_politica"
)

type FuerzaPoliticaInterface func() ([]fuerza_politica.FuerzaPolitica, error)

func FuerzaPoliticaGetService() FuerzaPoliticaInterface {
	return func() ([]fuerza_politica.FuerzaPolitica, error) {
		return fuerza_politica.GetFuerzasService()
	}
}

func FuerzaPoliticaGetServiceID(id int) FuerzaPoliticaInterface {
	return func() ([]fuerza_politica.FuerzaPolitica, error) {
		return fuerza_politica.GetFuerzasServiceID(id)
	}
}

func FuerzaPoliticaPostService(data interface{}) error {
	return fuerza_politica.PostFuerzaService(data)
}

func FuerzaPoliticaPutService(data interface{}) error {
	return fuerza_politica.PutFuerzaService(data)
}

func FuerzaPoliticaDeleteService(id int) error {
	return fuerza_politica.DeleteFuerzaService(id)
}