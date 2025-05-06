package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetHorario() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.HorarioGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetHorarioID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.HorarioGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostHorario(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.HorarioPostService(data)
	})
}

func PutHorario(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.HorarioPutService(data)
	})
}

func DeleteHorario(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.HorarioDeleteService(id)
	})
}