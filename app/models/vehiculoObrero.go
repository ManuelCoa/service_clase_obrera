package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetVehiculoObrero() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.VehiculoObreroGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetVehiculoObreroID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.VehiculoObreroGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostVehiculoObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.VehiculoObreroPostService(data)
	})
}

func PutVehiculoObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.VehiculoObreroPutService(data)
	})
}

func DeleteVehiculoObrero(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.VehiculoObreroDeleteService(id)
	})
}