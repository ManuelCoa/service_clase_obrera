package interfaz

import (
	"claseobrera/domains/object/comuna"
)

type ComunaInterface func() ([]comuna.Comuna, error)

func ComunaGetService() ComunaInterface {
	return func() ([]comuna.Comuna, error) {
		return comuna.GetComunaService()
	}
}

func ComunaGetServiceID(id int) ComunaInterface {
	return func() ([]comuna.Comuna, error) {
		return comuna.GetComunaServiceID(id)
	}
}

func ComunaPostService(data interface{}) error {
	return comuna.PostComunaService(data)
}

func ComunaPutService(data interface{}) error {
	return comuna.PutComunaService(data)
}

func ComunaDeleteService(id int) error {
	return comuna.DeleteComunaService(id)
}
