package interfaz

import (
	"claseobrera/domains/object/cargos_onapre"
)

type CargoOnapreInterface func() ([]cargos_onapre.CargosOnapre, error)

func CargoOnapreGetService() CargoOnapreInterface {
	return func() ([]cargos_onapre.CargosOnapre, error) {
		return cargos_onapre.GetCargoOnapreService()
	}
}

func CargoOnapreGetServiceID(id int) CargoOnapreInterface {
	return func() ([]cargos_onapre.CargosOnapre, error) {
		return cargos_onapre.GetCargoOnapreServiceID(id)
	}
}

func CargoOnaprePostService(data interface{}) error {
	return cargos_onapre.PostCargoOnapreService(data)
}

func CargoOnaprePutService(data interface{}) error {
	return cargos_onapre.PutCargoOnapreService(data)
}

func CargoOnapreDeleteService(id int) error {
	return cargos_onapre.DeleteCargoOnapreService(id)
}
