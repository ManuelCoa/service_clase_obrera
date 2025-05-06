package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoSueloVivienda() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoSueloViviendaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoSueloViviendaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoSueloViviendaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoSueloVivienda(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoSueloViviendaPostService(data)
	})
}

func PutTipoSueloVivienda(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoSueloViviendaPutService(data)
	})
}

func DeleteTipoSueloVivienda(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoSueloViviendaDeleteService(id)
	})
}