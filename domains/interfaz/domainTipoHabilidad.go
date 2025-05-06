package interfaz

import (
	"claseobrera/domains/object/tipo_habilidad"
)

type TipoHabilidadInterface func() ([]tipo_habilidad.TipoHabilidad, error)

func TipoHabilidadGetService() TipoHabilidadInterface {
	return func() ([]tipo_habilidad.TipoHabilidad, error) {
		return tipo_habilidad.GetTiposHabilidadService()
	}
}

func TipoHabilidadGetServiceID(id int) TipoHabilidadInterface {
	return func() ([]tipo_habilidad.TipoHabilidad, error) {
		return tipo_habilidad.GetTiposHabilidadServiceID(id)
	}
}

func TipoHabilidadPostService(data interface{}) error {
	return tipo_habilidad.PostTipoHabilidadService(data)
}

func TipoHabilidadPutService(data interface{}) error {
	return tipo_habilidad.PutTipoHabilidadService(data)
}

func TipoHabilidadDeleteService(id int) error {
	return tipo_habilidad.DeleteTipoHabilidadService(id)
}