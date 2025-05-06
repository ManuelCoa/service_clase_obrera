package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetAsistencias() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AsistenciaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetAsistenciaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AsistenciaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostAsistencia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaPostService(data)
	})
}

func PutAsistencia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaPutService(data)
	})
}

func DeleteAsistencia(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaDeleteService(id)
	})
}