package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetDireccionesObrero() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DireccionObreroGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDireccionObreroID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DireccionObreroGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}


func GetDireccionObreroInstitucionID(id_institucion int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DireccionObreroGetServiceInstitucionID(id_institucion)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostDireccionObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DireccionObreroPostService(data)
	})
}

func PutDireccionObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DireccionObreroPutService(data)
	})
}

func DeleteDireccionObrero(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DireccionObreroDeleteService(id)
	})
}