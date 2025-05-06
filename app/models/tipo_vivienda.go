package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoVivienda() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoViviendaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoViviendaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoViviendaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoVivienda(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoViviendaPostService(data)
	})
}

func PutTipoVivienda(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoViviendaPutService(data)
	})
}

func DeleteTipoVivienda(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoViviendaDeleteService(id)
	})
}