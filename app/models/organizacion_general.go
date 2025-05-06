package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetOrganizacionGeneral() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		org := interfaz.OrganizacionGeneralService()
		result, err := org()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetOrganizacionGeneralID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		orgID := interfaz.OrganizacionGeneralIDService(id)
		result, err := orgID()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostOrganizacionGeneral(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionGeneralPostService(data)
	})
}

func PutOrganizacionGeneral(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionGeneralPutService(data)
	})
}

func DeleteOrganizacionGeneralID(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.OrganizacionGeneralDeleteService(id)
	})
}
