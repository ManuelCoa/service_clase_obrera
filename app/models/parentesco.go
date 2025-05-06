package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetParentesco() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ParentescoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetParentescoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ParentescoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ParentescoPostService(data)
	})
}

func PutParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ParentescoPutService(data)
	})
}

func DeleteParentesco(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ParentescoDeleteService(id)
	})
}
