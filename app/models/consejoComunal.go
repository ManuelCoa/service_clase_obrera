package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetConsejoComunal() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ConsejoComunalGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetConsejoComunalID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ConsejoComunalGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostConsejoComunal(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ConsejoComunalPostService(data)
	})
}

func PutConsejoComunal(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ConsejoComunalPutService(data)
	})
}

func DeleteConsejoComunal(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ConsejoComunalDeleteService(id)
	})
}
