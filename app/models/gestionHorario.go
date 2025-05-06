package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetGestionHorarios() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GestionHorarioGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetGestionHorarioID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GestionHorarioGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostGestionHorario(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionHorarioPostService(data)
	})
}

func PutGestionHorario(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionHorarioPutService(data)
	})
}

func DeleteGestionHorario(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GestionHorarioDeleteService(id)
	})
}