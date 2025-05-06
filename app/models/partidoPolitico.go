package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetPartidoPolitico() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.PartidoPoliticoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetPartidoPoliticoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.PartidoPoliticoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostPartidoPolitico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PartidoPoliticoPostService(data)
	})
}

func PutPartidoPolitico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PartidoPoliticoPutService(data)
	})
}

func DeletePartidoPolitico(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PartidoPoliticoDeleteService(id)
	})
}
