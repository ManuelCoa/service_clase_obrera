package interfaz

import (
	org "claseobrera/domains/object/organizacion_estructura"
)

type OrganizacionEstructuraInterface func() ([]org.OrganizacionEstructura, error)

func OrganizacionEstructuraService() OrganizacionEstructuraInterface {
	return func() ([]org.OrganizacionEstructura, error) {
		return org.GetOrganizacionEstructuraService()
	}
}

func OrganizacionEstructuraIDService(id int) OrganizacionEstructuraInterface {
	return func() ([]org.OrganizacionEstructura, error) {
		return org.GetOrganizacionEstructuraIDService(id)
	}
}

func OrganizacionEstructuraPostService(data interface{}) error {
	return org.PostOrganizacionEstructuraService(data)
}

func OrganizacionEstructuraPutService(data interface{}) error {
	return org.PutOrganizacionEstructuraService(data)
}

func OrganizacionEstructuraDeleteService(id int) error {
	return org.DeleteOrganizacionEstructuraService(id)
}
