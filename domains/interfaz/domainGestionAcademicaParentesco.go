package interfaz

import (
	academica_parentesco "claseobrera/domains/object/gestion_academica_parentesco"
)

type AcademicaParentescoInterface func() ([]academica_parentesco.AcademicaParentesco, error)

func AcademicaParentescoGetService() AcademicaParentescoInterface {
	return func() ([]academica_parentesco.AcademicaParentesco, error) {
		return academica_parentesco.GetAcademicasParentescoService()
	}
}

func AcademicaParentescoGetServiceID(id int) AcademicaParentescoInterface {
	return func() ([]academica_parentesco.AcademicaParentesco, error) {
		return academica_parentesco.GetAcademicaParentescoServiceID(id)
	}
}

func AcademicaParentescoPostService(data interface{}) error {
	return academica_parentesco.PostAcademicaParentescoService(data)
}

func AcademicaParentescoPutService(data interface{}) error {
	return academica_parentesco.PutAcademicaParentescoService(data)
}

func AcademicaParentescoDeleteService(id int) error {
	return academica_parentesco.DeleteAcademicaParentescoService(id)
}