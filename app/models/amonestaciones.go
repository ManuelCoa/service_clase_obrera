package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetAmonestaciones() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AmonestacionesGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetAmonestacionID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AmonestacionesGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostAmonestacion(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AmonestacionesPostService(data)
	})
}

func PutAmonestacion(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AmonestacionesPutService(data)
	})
}

func DeleteAmonestacion(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AmonestacionesDeleteService(id)
	})
}