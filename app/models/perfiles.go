package models

import (
	"claseobrera/app/repository"
	"claseobrera/sessions/interfaz"
)

func GetPerfiles() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		org := interfaz.PerfilesGetService()
		result, err := org()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostPerfiles(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PerfilesPostService(data)
	})
}

func DeletePerfilesID(rolesID, permisosID int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PerfilesDeleteService(rolesID, permisosID)
	})
}
