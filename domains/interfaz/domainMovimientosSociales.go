package interfaz

import (
	"claseobrera/domains/object/movimientos_sociales"
)

type MovimientosSocialesInterface func() ([]movimientos_sociales.MovimientosSociales, error)

func MovimientosSocialesGetService() MovimientosSocialesInterface {
	return func() ([]movimientos_sociales.MovimientosSociales, error) {
		return movimientos_sociales.GetMovimientosSocialesService()
	}
}

func MovimientosSocialesGetServiceID(id int) MovimientosSocialesInterface {
	return func() ([]movimientos_sociales.MovimientosSociales, error) {
		return movimientos_sociales.GetMovimientosSocialesServiceID(id)
	}
}

func MovimientosSocialesPostService(data interface{}) error {
	return movimientos_sociales.PostMovimientosSocialesService(data)
}

func MovimientosSocialesPutService(data interface{}) error {
	return movimientos_sociales.PutMovimientosSocialesService(data)
}

func MovimientosSocialesDeleteService(id int) error {
	return movimientos_sociales.DeleteMovimientosSocialesService(id)
}