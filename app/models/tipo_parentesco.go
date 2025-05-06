package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoParentesco() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoParentescoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoParentescoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoParentescoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoParentescoPostService(data)
	})
}

func PutTipoParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoParentescoPutService(data)
	})
}

func DeleteTipoParentesco(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoParentescoDeleteService(id)
	})
}