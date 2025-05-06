package interfaz

import (
	"claseobrera/domains/object/solicitud_egreso"
)

type SolicitudEgresoInterface func() ([]solicitud_egreso.SolicitudEgreso, error)

func SolicitudEgresoGetService() SolicitudEgresoInterface {
	return func() ([]solicitud_egreso.SolicitudEgreso, error) {
		return solicitud_egreso.GetSolicitudEgresoService()
	}
}

func SolicitudEgresoGetServiceID(id int) SolicitudEgresoInterface {
	return func() ([]solicitud_egreso.SolicitudEgreso, error) {
		return solicitud_egreso.GetSolicitudEgresoServiceID(id)
	}
}

func SolicitudEgresoPostService(data interface{}) error {
	return solicitud_egreso.PostSolicitudEgresoService(data)
}

func SolicitudEgresoPutService(data interface{}) error {
	return solicitud_egreso.PutSolicitudEgresoService(data)
}

func SolicitudEgresoDeleteService(id int) error {
	return solicitud_egreso.DeleteSolicitudEgresoService(id)
}