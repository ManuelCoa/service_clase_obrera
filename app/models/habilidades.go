package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetHabilidades() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.HabilidadesGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetHabilidadID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.HabilidadesGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostHabilidad(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.HabilidadesPostService(data)
	})
}

func PutHabilidad(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.HabilidadesPutService(data)
	})
}

func DeleteHabilidad(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.HabilidadesDeleteService(id)
	})
}