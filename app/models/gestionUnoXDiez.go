package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetGestionUnoxDiez() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GestionUnoxDiezGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetGestionUnoxDiezID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GestionUnoxDiezGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostGestionUnoxDiez(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionUnoxDiezPostService(data)
	})
}

func PutGestionUnoxDiez(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionUnoxDiezPutService(data)
	})
}

func DeleteGestionUnoxDiez(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionUnoxDiezDeleteService(id)
	})
}