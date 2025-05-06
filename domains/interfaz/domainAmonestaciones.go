package interfaz

import (
	"claseobrera/domains/object/amonestaciones"
)

type AmonestacionesInterface func() ([]amonestaciones.Amonestaciones, error)

func AmonestacionesGetService() AmonestacionesInterface {
	return func() ([]amonestaciones.Amonestaciones, error) {
		return amonestaciones.GetAmonestacionesService()
	}
}

func AmonestacionesGetServiceID(id int) AmonestacionesInterface {
	return func() ([]amonestaciones.Amonestaciones, error) {
		return amonestaciones.GetAmonestacionServiceID(id)
	}
}

func AmonestacionesPostService(data interface{}) error {
	return amonestaciones.PostAmonestacionService(data)
}

func AmonestacionesPutService(data interface{}) error {
	return amonestaciones.PutAmonestacionService(data)
}

func AmonestacionesDeleteService(id int) error {
	return amonestaciones.DeleteAmonestacionService(id)
}