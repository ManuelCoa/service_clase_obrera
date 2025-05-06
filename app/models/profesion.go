package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetProfesion() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ProfesionGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetProfesionID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ProfesionGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostProfesion(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ProfesionPostService(data)
	})
}

func PutProfesion(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ProfesionPutService(data)
	})
}

func DeleteProfesion(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ProfesionDeleteService(id)
	})
}
