package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetGenero() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GeneroGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetGeneroID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GeneroGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostGenero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GeneroPostService(data)
	})
}

func PutGenero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GeneroPutService(data)
	})
}

func DeleteGenero(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GeneroDeleteService(id)
	})
}
