package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoEnfermedad() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoEnfermedadGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoEnfermedadID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoEnfermedadGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoEnfermedad(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoEnfermedadPostService(data)
	})
}

func PutTipoEnfermedad(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoEnfermedadPutService(data)
	})
}

func DeleteTipoEnfermedad(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoEnfermedadDeleteService(id)
	})
}