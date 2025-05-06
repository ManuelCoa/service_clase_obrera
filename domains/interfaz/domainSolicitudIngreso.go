package interfaz

import (
	"claseobrera/domains/object/solicitud_ingreso"
)

type SolicitudIngresoInterface func() ([]solicitud_ingreso.SolicitudIngreso, error)

func SolicitudIngresoGetService() SolicitudIngresoInterface {
	return func() ([]solicitud_ingreso.SolicitudIngreso, error) {
		return solicitud_ingreso.GetSolicitudIngresoService()
	}
}

func SolicitudIngresoGetServiceID(id int) SolicitudIngresoInterface {
	return func() ([]solicitud_ingreso.SolicitudIngreso, error) {
		return solicitud_ingreso.GetSolicitudIngresoServiceID(id)
	}
}

func SolicitudIngresoPostService(data interface{}) error {
	return solicitud_ingreso.PostSolicitudIngresoService(data)
}

func SolicitudIngresoPutService(data interface{}) error {
	return solicitud_ingreso.PutSolicitudIngresoService(data)
}

func SolicitudIngresoDeleteService(id int) error {
	return solicitud_ingreso.DeleteSolicitudIngresoService(id)
}