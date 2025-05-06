package interfaz

import (
	"claseobrera/domains/object/institucion_extracademica"
)

type InstitucionExtracademicaInterface func() ([]institucion_extracademica.InstitucionExtracademica, error)

func InstitucionExtracademicaGetService() InstitucionExtracademicaInterface {
	return func() ([]institucion_extracademica.InstitucionExtracademica, error) {
		return institucion_extracademica.GetInstitucionesService()
	}
}

func InstitucionExtracademicaGetServiceID(id int) InstitucionExtracademicaInterface {
	return func() ([]institucion_extracademica.InstitucionExtracademica, error) {
		return institucion_extracademica.GetInstitucionesServiceID(id)
	}
}

func InstitucionExtracademicaPostService(data interface{}) error {
	return institucion_extracademica.PostInstitucionService(data)
}

func InstitucionExtracademicaPutService(data interface{}) error {
	return institucion_extracademica.PutInstitucionService(data)
}

func InstitucionExtracademicaDeleteService(id int) error {
	return institucion_extracademica.DeleteInstitucionService(id)
}