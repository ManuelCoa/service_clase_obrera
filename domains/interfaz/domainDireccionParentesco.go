package interfaz

import (
	direccionparentesco "claseobrera/domains/object/gestion_direccion_parentesco"
)

type DireccionParentescoInterface func() ([]direccionparentesco.DireccionParentesco, error)

func DireccionParentescoGetService() DireccionParentescoInterface {
	return func() ([]direccionparentesco.DireccionParentesco, error) {
		return direccionparentesco.GetDireccionParentescoService()
	}
}

func DireccionParentescoGetServiceID(id int) DireccionParentescoInterface {
	return func() ([]direccionparentesco.DireccionParentesco, error) {
		return direccionparentesco.GetDireccionParentescoServiceID(id)
	}
}

func DireccionParentescoPostService(data interface{}) error {
	return direccionparentesco.PostDireccionParentescoService(data)
}

func DireccionParentescoPutService(data interface{}) error {
	return direccionparentesco.PutDireccionParentescoService(data)
}

func DireccionParentescoDeleteService(id int) error {
	return direccionparentesco.DeleteDireccionParentescoService(id)
}