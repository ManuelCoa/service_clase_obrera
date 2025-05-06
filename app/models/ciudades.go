package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetCiudad() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CiudadGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetCiudadID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CiudadGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostCiudad(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CiudadPostService(data)
	})
}

func PutCiudad(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CiudadPutService(data)
	})
}

func DeleteCiudad(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CiudadDeleteService(id)
	})
}
