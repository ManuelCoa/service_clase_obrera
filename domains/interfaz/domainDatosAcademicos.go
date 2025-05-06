package interfaz

import (
	datosacademicos "claseobrera/domains/object/datos_academicos"
)

type DatosAcademicosInterface func() ([]datosacademicos.DatosAcademicos, error)

func DatosAcademicosGetService() DatosAcademicosInterface {
	return func() ([]datosacademicos.DatosAcademicos, error) {
		return datosacademicos.GetDatosAcademicosService()
	}
}

func DatosAcademicosGetServiceID(id int) DatosAcademicosInterface {
	return func() ([]datosacademicos.DatosAcademicos, error) {
		return datosacademicos.GetDatosAcademicosServiceID(id)
	}
}

func DatosAcademicosPorObreroService(cedula int) DatosAcademicosInterface {
	return func() ([]datosacademicos.DatosAcademicos, error) {
		return datosacademicos.GetDatosAcademicosPorObreroService(cedula)
	}
}

func DatosAcademicosPostService(data interface{}) error {
	return datosacademicos.PostDatosAcademicosService(data)
}

func DatosAcademicosPutService(data interface{}) error {
	return datosacademicos.PutDatosAcademicosService(data)
}

func DatosAcademicosDeleteService(id int) error {
	return datosacademicos.DeleteDatosAcademicosService(id)
}