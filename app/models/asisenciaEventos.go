package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetAsistenciaEvento() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AsistenciaEventoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetAsistenciaEventoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AsistenciaEventoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostAsistenciaEvento(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaEventoPostService(data)
	})
}

func PutAsistenciaEvento(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaEventoPutService(data)
	})
}

func DeleteAsistenciaEvento(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaEventoDeleteService(id)
	})
}