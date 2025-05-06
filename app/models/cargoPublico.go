package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetCargo() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CargoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetCargoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CargoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostCargo(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CargoPostService(data)
	})
}

func PutCargo(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CargoPutService(data)
	})
}

func DeleteCargo(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CargoDeleteService(id)
	})
}