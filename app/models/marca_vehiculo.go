package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetMarcaVehiculo() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.MarcaVehiculoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetMarcaVehiculoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.MarcaVehiculoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostMarcaVehiculo(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MarcaVehiculoPostService(data)
	})
}

func PutMarcaVehiculo(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MarcaVehiculoPutService(data)
	})
}

func DeleteMarcaVehiculo(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MarcaVehiculoDeleteService(id)
	})
}