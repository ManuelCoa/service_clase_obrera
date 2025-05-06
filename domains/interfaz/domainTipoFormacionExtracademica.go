package interfaz

import (
	"claseobrera/domains/object/tipo_formacion_extracademica"
)

type TipoFormacionExtracademicaInterface func() ([]tipo_formacion_extracademica.TipoFormacionExtracademica, error)

func TipoFormacionExtracademicaGetService() TipoFormacionExtracademicaInterface {
	return func() ([]tipo_formacion_extracademica.TipoFormacionExtracademica, error) {
		return tipo_formacion_extracademica.GetTiposFormacionService()
	}
}

func TipoFormacionExtracademicaGetServiceID(id int) TipoFormacionExtracademicaInterface {
	return func() ([]tipo_formacion_extracademica.TipoFormacionExtracademica, error) {
		return tipo_formacion_extracademica.GetTiposFormacionServiceID(id)
	}
}

func TipoFormacionExtracademicaPostService(data interface{}) error {
	return tipo_formacion_extracademica.PostTipoFormacionService(data)
}

func TipoFormacionExtracademicaPutService(data interface{}) error {
	return tipo_formacion_extracademica.PutTipoFormacionService(data)
}

func TipoFormacionExtracademicaDeleteService(id int) error {
	return tipo_formacion_extracademica.DeleteTipoFormacionService(id)
}