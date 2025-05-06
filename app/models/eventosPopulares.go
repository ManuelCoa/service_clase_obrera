package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetEventoPopular() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.EventoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetEventoPopularID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.EventoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostEventoPopular(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EventoPostService(data)
	})
}

func PutEventoPopular(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EventoPutService(data)
	})
}

func DeleteEventoPopular(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EventoDeleteService(id)
	})
}