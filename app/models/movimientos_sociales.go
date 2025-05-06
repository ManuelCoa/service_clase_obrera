package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetMovimientosSociales() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GeneroGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetMovimientosSocialesID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.MovimientosSocialesGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostMovimientosSociales(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MovimientosSocialesPostService(data)
	})
}

func PutMovimientosSociales(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MovimientosSocialesPutService(data)
	})
}

func DeleteMovimientosSociales(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MovimientosSocialesDeleteService(id)
	})
}
