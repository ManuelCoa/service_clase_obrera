package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetComuna() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ComunaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetComunaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ComunaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostComuna(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ComunaPostService(data)
	})
}

func PutComuna(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CiudadPutService(data)
	})
}

func DeleteComuna(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ComunaDeleteService(id)
	})
}

