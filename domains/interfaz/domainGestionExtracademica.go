package interfaz

import (
	gestionextracademica "claseobrera/domains/object/gestion_extracademica"
)

type GestionExtracademicaInterface func() ([]gestionextracademica.GestionExtracademica, error)

func GestionExtracademicaGetService() GestionExtracademicaInterface {
	return func() ([]gestionextracademica.GestionExtracademica, error) {
		return gestionextracademica.GetGestionExtracademicaService()
	}
}

func GestionExtracademicaGetServiceID(id int) GestionExtracademicaInterface {
	return func() ([]gestionextracademica.GestionExtracademica, error) {
		return gestionextracademica.GetGestionExtracademicaServiceID(id)
	}
}

func GestionExtracademicaPostService(data interface{}) error {
	return gestionextracademica.PostGestionExtracademicaService(data)
}

func GestionExtracademicaPutService(data interface{}) error {
	return gestionextracademica.PutGestionExtracademicaService(data)
}

func GestionExtracademicaDeleteService(id int) error {
	return gestionextracademica.DeleteGestionExtracademicaService(id)
}