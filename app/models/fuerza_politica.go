package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetFuerzaPolitica() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.FuerzaPoliticaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetFuerzaPoliticaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.FuerzaPoliticaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostFuerzaPolitica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.FuerzaPoliticaPostService(data)
	})
}

func PutFuerzaPolitica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.FuerzaPoliticaPutService(data)
	})
}

func DeleteFuerzaPolitica(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.FuerzaPoliticaDeleteService(id)
	})
}