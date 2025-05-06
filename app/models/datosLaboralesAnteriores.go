package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetDatosLaborales() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DatosLaboralesGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDatosLaboralesID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DatosLaboralesGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDatosLaboralesPorObrero(cedula int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DatosLaboralesPorObreroService(cedula)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostDatosLaborales(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DatosLaboralesPostService(data)
	})
}

func PutDatosLaborales(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DatosLaboralesPutService(data)
	})
}

func DeleteDatosLaborales(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DatosLaboralesDeleteService(id)
	})
}