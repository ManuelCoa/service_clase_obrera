package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetVivienda() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ViviendaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetViviendaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ViviendaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostVivienda(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ViviendaPostService(data)
	})
}

func PutVivienda(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ViviendaPutService(data)
	})
}

func DeleteVivienda(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ViviendaDeleteService(id)
	})
}