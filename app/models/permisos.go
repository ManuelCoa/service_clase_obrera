package models

import (
	"claseobrera/app/repository"
	"claseobrera/sessions/interfaz"
)

func GetPermisos() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		org := interfaz.PermisosGetService()
		result, err := org()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetPermisosID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		orgID := interfaz.PermisosGetServiceID(id)
		result, err := orgID()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostPermisos(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PermisosPostService(data)
	})
}

func PutPermisos(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PermisosPutService(data)
	})
}

func DeletePermisosID(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PermisosDeleteService(id)
	})
}
