package interfaz

import (
	"claseobrera/sessions/object/permisos"
)

type PermisosInterface func() ([]permisos.Permisos, error)

func PermisosGetService() PermisosInterface {
	return func() ([]permisos.Permisos, error) {
		return permisos.GetPermisosService()
	}
}

func PermisosGetServiceID(id int) PermisosInterface {
	return func() ([]permisos.Permisos, error) {
		return permisos.GetPermisosIDService(id)
	}
}

func PermisosPostService(data interface{}) error {
	return permisos.PostPermisosService(data)
}

func PermisosPutService(data interface{}) error {
	return permisos.PutPermisosService(data)
}

func PermisosDeleteService(id int) error {
	return permisos.DeletePermisosService(id)
}
