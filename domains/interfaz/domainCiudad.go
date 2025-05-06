package interfaz

import (
	"claseobrera/domains/object/ciudades"
)

type CiudadInterface func() ([]ciudades.Ciudad, error)

func CiudadGetService() CiudadInterface {
	return func() ([]ciudades.Ciudad, error) {
		return ciudades.GetCiudadService()
	}
}

func CiudadGetServiceID(id int) CiudadInterface {
	return func() ([]ciudades.Ciudad, error) {
		return ciudades.GetCiudadServiceID(id)
	}
}

func CiudadPostService(data interface{}) error {
	return ciudades.PostCiudadService(data)
}

func CiudadPutService(data interface{}) error {
	return ciudades.PutCiudadService(data)
}

func CiudadDeleteService(id int) error {
	return ciudades.DeleteCiudadService(id)
}
