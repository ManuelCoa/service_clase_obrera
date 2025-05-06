package interfaz

import (
	"claseobrera/domains/object/salud_parentesco"
)

type SaludParentescoInterface func() ([]salud_parentesco.SaludParentesco, error)

func SaludParentescoGetService() SaludParentescoInterface {
	return func() ([]salud_parentesco.SaludParentesco, error) {
		return salud_parentesco.GetSaludParentescoService()
	}
}

func SaludParentescoGetServiceID(id int) SaludParentescoInterface {
	return func() ([]salud_parentesco.SaludParentesco, error) {
		return salud_parentesco.GetSaludParentescoServiceID(id)
	}
}

func SaludParentescoPostService(data interface{}) error {
	return salud_parentesco.PostSaludParentescoService(data)
}

func SaludParentescoPutService(data interface{}) error {
	return salud_parentesco.PutSaludParentescoService(data)
}

func SaludParentescoDeleteService(id int) error {
	return salud_parentesco.DeleteSaludParentescoService(id)
}