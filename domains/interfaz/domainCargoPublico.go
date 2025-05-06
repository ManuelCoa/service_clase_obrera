package interfaz

import (
	"claseobrera/domains/object/cargo_publico"
)

type CargoInterface func() ([]cargo_publico.CargoPublico, error)

func CargoGetService() CargoInterface {
	return func() ([]cargo_publico.CargoPublico, error) {
		return cargo_publico.GetCargoService()
	}
}

func CargoGetServiceID(id int) CargoInterface {
	return func() ([]cargo_publico.CargoPublico, error) {
		return cargo_publico.GetCargoServiceID(id)
	}
}

func CargoPostService(data interface{}) error {
	return cargo_publico.PostCargoService(data)
}

func CargoPutService(data interface{}) error {
	return cargo_publico.PutCargoService(data)
}

func CargoDeleteService(id int) error {
	return cargo_publico.DeleteCargoService(id)
}