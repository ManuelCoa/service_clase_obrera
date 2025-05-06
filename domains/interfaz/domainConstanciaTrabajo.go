package interfaz

import (
	constancias "claseobrera/domains/object/constancia_trabajo"
)

type ConstanciaInterface func() ([]constancias.ConstanciaTrabajo, error)

func ConstanciaGetService() ConstanciaInterface {
	return func() ([]constancias.ConstanciaTrabajo, error) {
		return constancias.GetConstanciasService()
	}
}

func ConstanciaGetServiceID(id int) ConstanciaInterface {
	return func() ([]constancias.ConstanciaTrabajo, error) {
		return constancias.GetConstanciaServiceID(id)
	}
}

func ConstanciaPostService(data interface{}) error {
	return constancias.PostConstanciaService(data)
}

func ConstanciaPutService(data interface{}) error {
	return constancias.PutConstanciaService(data)
}

func ConstanciaDeleteService(id int) error {
	return constancias.DeleteConstanciaService(id)
}