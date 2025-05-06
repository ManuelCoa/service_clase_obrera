package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetDepartamentos() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DepartamentosGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDepartamentosID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DepartamentosGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostDepartamentos(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DepartamentosPostService(data)
	})
}

func PutDepartamentos(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DepartamentosPutService(data)
	})
}

func DeleteDepartamentos(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DepartamentosDeleteService(id)
	})
}