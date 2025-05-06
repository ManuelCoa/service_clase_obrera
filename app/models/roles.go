package models

import (
	"claseobrera/app/repository"
	"claseobrera/sessions/interfaz"
)

func GetRoles() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		org := interfaz.RolesGetService()
		result, err := org()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetRolesID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		orgID := interfaz.RolesGetServiceID(id)
		result, err := orgID()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostRoles(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.RolesPostService(data)
	})
}

func PutRoles(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.RolesPutService(data)
	})
}

func DeleteRolesID(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.RolesDeleteService(id)
	})
}
