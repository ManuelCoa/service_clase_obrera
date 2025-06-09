package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTestPiel() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TestPielGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTestPielID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TestPielGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTestPiel(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TestPielPostService(data)
	})
}

func PutTestPiel(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TestPielPutService(data)
	})
}

func DeleteTestPiel(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TestPielDeleteService(id)
	})
}