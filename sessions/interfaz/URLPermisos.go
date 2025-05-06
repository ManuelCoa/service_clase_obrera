package interfaz

import (
	"claseobrera/sessions/object/permisosURL"
)

type PermisosURLInterface func() ([]permisosURL.PermisosURL, error)

func PermisosURLGetServiceID(id int) PermisosURLInterface {
	return func() ([]permisosURL.PermisosURL, error) {
		return permisosURL.GetPermissionsByRoleID(id)
	}
}
