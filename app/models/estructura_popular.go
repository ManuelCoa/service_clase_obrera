package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetEstructuraPopular() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.EstructuraPopularGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetEstructuraPopularID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.EstructuraPopularGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostEstructuraPopular(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstructuraPopularPostService(data)
	})
}

func PutEstructuraPopular(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstructuraPopularPutService(data)
	})
}

func DeleteEstructuraPopular(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.EstructuraPopularDeleteService(id)
	})
}

