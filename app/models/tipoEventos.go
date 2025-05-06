package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoEvento() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoEventoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoEventoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoEventoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoEvento(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoEventoPostService(data)
	})
}

func PutTipoEvento(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoEventoPutService(data)
	})
}

func DeleteTipoEvento(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoEventoDeleteService(id)
	})
}