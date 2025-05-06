package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetCargoOnapre() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CargoOnapreGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetCargoOnapreID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CargoOnapreGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostCargoOnapre(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CargoOnaprePostService(data)
	})
}

func PutCargoOnapre(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CargoOnaprePutService(data)
	})
}

func DeleteCargoOnapre(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CargoOnapreDeleteService(id)
	})
}
