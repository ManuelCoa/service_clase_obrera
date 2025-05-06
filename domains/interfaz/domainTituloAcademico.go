package interfaz

import (
	"claseobrera/domains/object/titulo_academico"
)

type TituloAcademicoInterface func() ([]titulo_academico.TituloAcademico, error)

func TituloAcademicoGetService() TituloAcademicoInterface {
	return func() ([]titulo_academico.TituloAcademico, error) {
		return titulo_academico.GetTituloAcademicoService()
	}
}

func TituloAcademicoGetServiceID(id int) TituloAcademicoInterface {
	return func() ([]titulo_academico.TituloAcademico, error) {
		return titulo_academico.GetTituloAcademicoServiceID(id)
	}
}

func TituloAcademicoPostService(data interface{}) error {
	return titulo_academico.PostTituloAcademicoService(data)
}

func TituloAcademicoPutService(data interface{}) error {
	return titulo_academico.PutTituloAcademicoService(data)
}

func TituloAcademicoDeleteService(id int) error {
	return titulo_academico.DeleteTituloAcademicoService(id)
}