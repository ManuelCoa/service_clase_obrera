package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetOrganizacionJerarquia() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		org := interfaz.OrganizacionJerarquiaService()
		result, err := org()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetOrganizacionJerarquiaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		orgID := interfaz.OrganizacionJerarquiaIDService(id)
		result, err := orgID()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostOrganizacionJerarquia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionJerarquiaPostService(data)
	})
}

func PutOrganizacionJerarquia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionJerarquiaPutService(data)
	})
}

func DeleteOrganizacionJerarquiaID(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionJerarquiaDeleteService(id)
	})
}
