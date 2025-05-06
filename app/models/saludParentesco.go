package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetSaludParentescos() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.SaludParentescoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetSaludParentescoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.SaludParentescoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostSaludParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SaludParentescoPostService(data)
	})
}

func PutSaludParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SaludParentescoPutService(data)
	})
}

func DeleteSaludParentesco(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SaludParentescoDeleteService(id)
	})
}