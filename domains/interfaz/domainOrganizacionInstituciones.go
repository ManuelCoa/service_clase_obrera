package interfaz

import (
	org "claseobrera/domains/object/organizacion_instituciones"
)

type InstitucionesInterface func() ([]org.Instituciones, error)

func InstitucionesService() InstitucionesInterface {
	return func() ([]org.Instituciones, error) {
		return org.GetInstitucionesService()
	}
}

func InstitucionesIDService(id int) InstitucionesInterface {
	return func() ([]org.Instituciones, error) {
		return org.GetInstitucionesIDService(id)
	}
}

func InstitucionesPostService(data interface{}) error {
	return org.PostInstitucionesService(data)
}

func InstitucionesPutService(data interface{}) error {
	return org.PutCategoriaOrganizacionService(data)
}

func InstitucionesDeleteService(id int) error {
	return org.DeleteInstitucionesService(id)
}
