package interfaz

import (
	datoslaborales "claseobrera/domains/object/datos_laborales_anteriores"
)

type DatosLaboralesInterface func() ([]datoslaborales.DatosLaboralesAnteriores, error)

func DatosLaboralesGetService() DatosLaboralesInterface {
	return func() ([]datoslaborales.DatosLaboralesAnteriores, error) {
		return datoslaborales.GetDatosLaboralesService()
	}
}

func DatosLaboralesGetServiceID(id int) DatosLaboralesInterface {
	return func() ([]datoslaborales.DatosLaboralesAnteriores, error) {
		return datoslaborales.GetDatosLaboralesServiceID(id)
	}
}

func DatosLaboralesPorObreroService(cedula int) DatosLaboralesInterface {
	return func() ([]datoslaborales.DatosLaboralesAnteriores, error) {
		return datoslaborales.GetDatosLaboralesPorObreroService(cedula)
	}
}

func DatosLaboralesPostService(data interface{}) error {
	return datoslaborales.PostDatosLaboralesService(data)
}

func DatosLaboralesPutService(data interface{}) error {
	return datoslaborales.PutDatosLaboralesService(data)
}

func DatosLaboralesDeleteService(id int) error {
	return datoslaborales.DeleteDatosLaboralesService(id)
}