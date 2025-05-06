package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetEstado() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.EstadoService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetEstadoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.EstadoIDService(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostEstado(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstadoPostService(data)
	})
}

func PutEstado(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstadoPutService(data)
	})
}

func DeleteEstadoID(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstadoDeleteService(id)
	})
}
