package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetParroquia() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ParroquiaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetParroquiaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ParroquiaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostParroquia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ParroquiaPostService(data)
	})
}

func PutParroquia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ParroquiaPutService(data)
	})
}

func DeleteParroquia(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ParroquiaDeleteService(id)
	})
}
