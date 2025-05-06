package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoTecho() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoTechoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoTechoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoTechoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoTecho(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoTechoPostService(data)
	})
}

func PutTipoTecho(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoTechoPutService(data)
	})
}

func DeleteTipoTecho(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoTechoDeleteService(id)
	})
}