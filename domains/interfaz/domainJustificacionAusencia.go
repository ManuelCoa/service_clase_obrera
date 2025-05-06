package interfaz

import (
	"claseobrera/domains/object/justificacion_ausencia"
)

type JustificacionAusenciaInterface func() ([]justificacion_ausencia.JustificacionAusencia, error)

func JustificacionAusenciaGetService() JustificacionAusenciaInterface {
	return func() ([]justificacion_ausencia.JustificacionAusencia, error) {
		return justificacion_ausencia.GetJustificacionesService()
	}
}

func JustificacionAusenciaGetServiceID(id int) JustificacionAusenciaInterface {
	return func() ([]justificacion_ausencia.JustificacionAusencia, error) {
		return justificacion_ausencia.GetJustificacionServiceID(id)
	}
}

func JustificacionAusenciaPostService(data interface{}) error {
	return justificacion_ausencia.PostJustificacionService(data)
}

func JustificacionAusenciaPutService(data interface{}) error {
	return justificacion_ausencia.PutJustificacionService(data)
}

func JustificacionAusenciaDeleteService(id int) error {
	return justificacion_ausencia.DeleteJustificacionService(id)
}