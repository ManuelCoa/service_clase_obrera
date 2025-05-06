package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoPared() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoParedGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoParedID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoParedGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoPared(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoParedPostService(data)
	})
}

func PutTipoPared(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoParedPutService(data)
	})
}

func DeleteTipoPared(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoParedDeleteService(id)
	})
}