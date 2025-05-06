package interfaz

import (
	"claseobrera/domains/object/asistencia"
)

type AsistenciaInterface func() ([]asistencia.Asistencia, error)

func AsistenciaGetService() AsistenciaInterface {
	return func() ([]asistencia.Asistencia, error) {
		return asistencia.GetAsistenciasService()
	}
}

func AsistenciaGetServiceID(id int) AsistenciaInterface {
	return func() ([]asistencia.Asistencia, error) {
		return asistencia.GetAsistenciaServiceID(id)
	}
}

func AsistenciaPostService(data interface{}) error {
	return asistencia.PostAsistenciaService(data)
}

func AsistenciaPutService(data interface{}) error {
	return asistencia.PutAsistenciaService(data)
}

func AsistenciaDeleteService(id int) error {
	return asistencia.DeleteAsistenciaService(id)
}