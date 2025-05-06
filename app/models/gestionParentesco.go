package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetGestionParentescos() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GestionParentescoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetGestionParentescoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GestionParentescoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostGestionParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionParentescoPostService(data)
	})
}

func PutGestionParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionParentescoPutService(data)
	})
}

func DeleteGestionParentesco(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionParentescoDeleteService(id)
	})
}