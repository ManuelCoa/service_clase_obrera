package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetDireccionParentesco() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DireccionParentescoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDireccionParentescoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DireccionParentescoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostDireccionParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DireccionParentescoPostService(data)
	})
}

func PutDireccionParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DireccionParentescoPutService(data)
	})
}

func DeleteDireccionParentesco(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DireccionParentescoDeleteService(id)
	})
}