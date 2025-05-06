package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoHabilidad() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoHabilidadGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoHabilidadID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoHabilidadGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoHabilidad(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoHabilidadPostService(data)
	})
}

func PutTipoHabilidad(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoHabilidadPutService(data)
	})
}

func DeleteTipoHabilidad(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoHabilidadDeleteService(id)
	})
}