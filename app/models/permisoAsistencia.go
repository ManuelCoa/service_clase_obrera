package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetPermisosAsistencia() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.PermisoAsistenciaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetPermisoAsistenciaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.PermisoAsistenciaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostPermisoAsistencia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PermisoAsistenciaPostService(data)
	})
}

func PutPermisoAsistencia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PermisoAsistenciaPutService(data)
	})
}

func DeletePermisoAsistencia(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PermisoAsistenciaDeleteService(id)
	})
}