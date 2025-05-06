package interfaz

import (
	"claseobrera/domains/object/tipo_enfermedad"
)

type TipoEnfermedadInterface func() ([]tipo_enfermedad.TipoEnfermedad, error)

func TipoEnfermedadGetService() TipoEnfermedadInterface {
	return func() ([]tipo_enfermedad.TipoEnfermedad, error) {
		return tipo_enfermedad.GetTiposEnfermedadService()
	}
}

func TipoEnfermedadGetServiceID(id int) TipoEnfermedadInterface {
	return func() ([]tipo_enfermedad.TipoEnfermedad, error) {
		return tipo_enfermedad.GetTiposEnfermedadServiceID(id)
	}
}

func TipoEnfermedadPostService(data interface{}) error {
	return tipo_enfermedad.PostTipoEnfermedadService(data)
}

func TipoEnfermedadPutService(data interface{}) error {
	return tipo_enfermedad.PutTipoEnfermedadService(data)
}

func TipoEnfermedadDeleteService(id int) error {
	return tipo_enfermedad.DeleteTipoEnfermedadService(id)
}