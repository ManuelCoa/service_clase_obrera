package interfaz

import (
	"claseobrera/domains/object/parroquias"
)

type ParroquiaInterface func() ([]parroquias.Parroquia, error)

func ParroquiaGetService() ParroquiaInterface {
	return func() ([]parroquias.Parroquia, error) {
		return parroquias.GetParroquiaService()
	}
}

func ParroquiaGetServiceID(id int) ParroquiaInterface {
	return func() ([]parroquias.Parroquia, error) {
		return parroquias.GetParroquiaServiceID(id)
	}
}

func ParroquiaPostService(data interface{}) error {
	return parroquias.PostParroquiaService(data)
}

func ParroquiaPutService(data interface{}) error {
	return parroquias.PutParroquiaService(data)
}

func ParroquiaDeleteService(id int) error {
	return parroquias.DeleteParroquiaService(id)
}
