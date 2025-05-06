package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetMisionesBase() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.MisionesBaseGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetMisionesBaseID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.MisionesBaseGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostMisionesBase(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MisionesBasePostService(data)
	})
}

func PutMisionesBase(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MisionesBasePutService(data)
	})
}

func DeleteMisionesBase(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MisionesBaseDeleteService(id)
	})
}