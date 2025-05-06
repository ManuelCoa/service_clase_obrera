package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetConstancias() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ConstanciaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetConstanciaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ConstanciaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostConstancia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ConstanciaPostService(data)
	})
}

func PutConstancia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ConstanciaPutService(data)
	})
}

func DeleteConstancia(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ConstanciaDeleteService(id)
	})
}