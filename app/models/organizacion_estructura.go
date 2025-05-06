package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetOrganizacionEstructura() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		org := interfaz.OrganizacionEstructuraService()
		result, err := org()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetOrganizacionEstructuraID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		orgID := interfaz.OrganizacionEstructuraIDService(id)
		result, err := orgID()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostOrganizacionEstructura(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionEstructuraPostService(data)
	})
}

func PutOrganizacionEstructura(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionEstructuraPutService(data)
	})
}

func DeleteOrganizacionEstructuraID(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionEstructuraDeleteService(id)
	})
}
