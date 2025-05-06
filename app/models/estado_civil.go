package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetEstadoCivil() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.EstadoCivilGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetEstadoCivilID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.EstadoCivilGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostEstadoCivil(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstadoCivilPostService(data)
	})
}

func PutEstadoCivil(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstadoCivilPutService(data)
	})
}

func DeleteEstadoCivil(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstadoCivilDeleteService(id)
	})
}
