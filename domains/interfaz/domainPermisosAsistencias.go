package interfaz

import (
	"claseobrera/domains/object/permisos_asistencias"
)

type PermisoAsistenciaInterface func() ([]permisos_asistencias.PermisoAsistencia, error)

func PermisoAsistenciaGetService() PermisoAsistenciaInterface {
	return func() ([]permisos_asistencias.PermisoAsistencia, error) {
		return permisos_asistencias.GetPermisoAsistenciaService()
	}
}

func PermisoAsistenciaGetServiceID(id int) PermisoAsistenciaInterface {
	return func() ([]permisos_asistencias.PermisoAsistencia, error) {
		return permisos_asistencias.GetPermisoAsistenciaServiceID(id)
	}
}

func PermisoAsistenciaPostService(data interface{}) error {
	return permisos_asistencias.PostPermisoAsistenciaService(data)
}

func PermisoAsistenciaPutService(data interface{}) error {
	return permisos_asistencias.PutPermisoAsistenciaService(data)
}

func PermisoAsistenciaDeleteService(id int) error {
	return permisos_asistencias.DeletePermisoAsistenciaService(id)
}