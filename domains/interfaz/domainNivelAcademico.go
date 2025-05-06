package interfaz

import (
	"claseobrera/domains/object/nivel_academico"
)

type NivelAcademicoInterface func() ([]nivel_academico.NivelAcademico, error)

func NivelAcademicoGetService() NivelAcademicoInterface {
	return func() ([]nivel_academico.NivelAcademico, error) {
		return nivel_academico.GetNivelAcademicoService()
	}
}

func NivelAcademicoGetServiceID(id int) NivelAcademicoInterface {
	return func() ([]nivel_academico.NivelAcademico, error) {
		return nivel_academico.GetNivelAcademicoServiceID(id)
	}
}

func NivelAcademicoPostService(data interface{}) error {
	return nivel_academico.PostNivelAcademicoService(data)
}

func NivelAcademicoPutService(data interface{}) error {
	return nivel_academico.PutNivelAcademicoService(data)
}

func NivelAcademicoDeleteService(id int) error {
	return nivel_academico.DeleteNivelAcademicoService(id)
}
