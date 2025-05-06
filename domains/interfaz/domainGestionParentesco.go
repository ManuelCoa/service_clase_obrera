package interfaz

import (
	"claseobrera/domains/object/gestion_parentesco"
)

type GestionParentescoInterface func() ([]gestion_parentesco.GestionParentesco, error)

func GestionParentescoGetService() GestionParentescoInterface {
	return func() ([]gestion_parentesco.GestionParentesco, error) {
		return gestion_parentesco.GetParentescoService()
	}
}

func GestionParentescoGetServiceID(id int) GestionParentescoInterface {
	return func() ([]gestion_parentesco.GestionParentesco, error) {
		return gestion_parentesco.GetParentescoServiceID(id)
	}
}

func GestionParentescoPostService(data interface{}) error {
	return gestion_parentesco.PostParentescoService(data)
}

func GestionParentescoPutService(data interface{}) error {
	return gestion_parentesco.PutParentescoService(data)
}

func GestionParentescoDeleteService(id int) error {
	return gestion_parentesco.DeleteParentescoService(id)
}