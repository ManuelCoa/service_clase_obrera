package interfaz

import (
	"claseobrera/domains/object/institucion_academica"
)

type InstitucionAcademicaInterface func() ([]institucion_academica.InstitucionAcademica, error)

func InstitucionAcademicaGetService() InstitucionAcademicaInterface {
	return func() ([]institucion_academica.InstitucionAcademica, error) {
		return institucion_academica.GetInstitucionesAcademicasService()
	}
}

func InstitucionAcademicaGetServiceID(id int) InstitucionAcademicaInterface {
	return func() ([]institucion_academica.InstitucionAcademica, error) {
		return institucion_academica.GetInstitucionAcademicaServiceID(id)
	}
}

func InstitucionAcademicaPostService(data interface{}) error {
	return institucion_academica.PostInstitucionAcademicaService(data)
}

func InstitucionAcademicaPutService(data interface{}) error {
	return institucion_academica.PutInstitucionAcademicaService(data)
}

func InstitucionAcademicaDeleteService(id int) error {
	return institucion_academica.DeleteInstitucionAcademicaService(id)
}