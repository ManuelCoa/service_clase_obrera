package interfaz

import (
	org "claseobrera/domains/object/organizacion_jerarquia"
)

type OrganizacionJerarquiaInterface func() ([]org.OrganizacionJerarquia, error)

func OrganizacionJerarquiaService() OrganizacionJerarquiaInterface {
	return func() ([]org.OrganizacionJerarquia, error) {
		return org.GetOrganizacionJerarquiaService()
	}
}

func OrganizacionJerarquiaIDService(id int) OrganizacionJerarquiaInterface {
	return func() ([]org.OrganizacionJerarquia, error) {
		return org.GetOrganizacionJerarquiaIDService(id)
	}
}

func OrganizacionJerarquiaPostService(data interface{}) error {
	return org.PostOrganizacionJerarquiaService(data)
}

func OrganizacionJerarquiaPutService(data interface{}) error {
	return org.PutOrganizacionJerarquiaService(data)
}

func OrganizacionJerarquiaDeleteService(id int) error {
	return org.DeleteOrganizacionJerarquiaService(id)
}
