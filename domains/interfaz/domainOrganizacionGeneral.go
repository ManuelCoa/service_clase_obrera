package interfaz

import (
	org "claseobrera/domains/object/organizacion_general"
)

type OrganizacionGeneralInterface func() ([]org.OrganizacionGeneral, error)

func OrganizacionGeneralService() OrganizacionGeneralInterface {
	return func() ([]org.OrganizacionGeneral, error) {
		return org.GetOrganizacionGeneralService()
	}
}

func OrganizacionGeneralIDService(id int) OrganizacionGeneralInterface {
	return func() ([]org.OrganizacionGeneral, error) {
		return org.GetOrganizacionGeneralIDService(id)
	}
}

func OrganizacionGeneralPostService(data interface{}) error {
	return org.PostOrganizacionGeneralService(data)
}

func OrganizacionGeneralPutService(data interface{}) error {
	return org.PutOrganizacionGeneralService(data)
}

func OrganizacionGeneralDeleteService(id int) error {
	return org.DeleteOrganizacionGeneralService(id)
}
