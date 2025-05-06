package interfaz

import (
	"claseobrera/domains/object/genero"
)

type GeneroInterface func() ([]genero.Genero, error)

func GeneroGetService() GeneroInterface {
	return func() ([]genero.Genero, error) {
		return genero.GetGeneroService()
	}
}

func GeneroGetServiceID(id int) GeneroInterface {
	return func() ([]genero.Genero, error) {
		return genero.GetGeneroServiceID(id)
	}
}

func GeneroPostService(data interface{}) error {
	return genero.PostGeneroService(data)
}

func GeneroPutService(data interface{}) error {
	return genero.PutGeneroService(data)
}

func GeneroDeleteService(id int) error {
	return genero.DeleteGeneroService(id)
}