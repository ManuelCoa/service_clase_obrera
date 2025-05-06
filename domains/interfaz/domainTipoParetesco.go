package interfaz

import (
	"claseobrera/domains/object/tipo_parentesco"
)

type TipoParentescoInterface func() ([]tipo_parentesco.TipoParentesco, error)

func TipoParentescoGetService() TipoParentescoInterface {
	return func() ([]tipo_parentesco.TipoParentesco, error) {
		return tipo_parentesco.GetTiposParentescoService()
	}
}

func TipoParentescoGetServiceID(id int) TipoParentescoInterface {
	return func() ([]tipo_parentesco.TipoParentesco, error) {
		return tipo_parentesco.GetTiposParentescoServiceID(id)
	}
}

func TipoParentescoPostService(data interface{}) error {
	return tipo_parentesco.PostTipoParentescoService(data)
}

func TipoParentescoPutService(data interface{}) error {
	return tipo_parentesco.PutTipoParentescoService(data)
}

func TipoParentescoDeleteService(id int) error {
	return tipo_parentesco.DeleteTipoParentescoService(id)
}