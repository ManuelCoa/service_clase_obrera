package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetGestionesExtracademicas() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GestionExtracademicaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetGestionExtracademicaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GestionExtracademicaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostGestionExtracademica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionExtracademicaPostService(data)
	})
}

func PutGestionExtracademica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionExtracademicaPutService(data)
	})
}

func DeleteGestionExtracademica(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionExtracademicaDeleteService(id)
	})
}