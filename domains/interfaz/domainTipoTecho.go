package interfaz

import (
	"claseobrera/domains/object/tipo_techo"
)

type TipoTechoInterface func() ([]tipo_techo.TipoTecho, error)

func TipoTechoGetService() TipoTechoInterface {
	return func() ([]tipo_techo.TipoTecho, error) {
		return tipo_techo.GetTiposTechoService()
	}
}

func TipoTechoGetServiceID(id int) TipoTechoInterface {
	return func() ([]tipo_techo.TipoTecho, error) {
		return tipo_techo.GetTiposTechoServiceID(id)
	}
}

func TipoTechoPostService(data interface{}) error {
	return tipo_techo.PostTipoTechoService(data)
}

func TipoTechoPutService(data interface{}) error {
	return tipo_techo.PutTipoTechoService(data)
}

func TipoTechoDeleteService(id int) error {
	return tipo_techo.DeleteTipoTechoService(id)
}