package interfaz

import (
	"claseobrera/sessions/object/roles"
)

type RolesInterface func() ([]roles.Roles, error)

func RolesGetService() RolesInterface {
	return func() ([]roles.Roles, error) {
		return roles.GetRolesService()
	}
}

func RolesGetServiceID(id int) RolesInterface {
	return func() ([]roles.Roles, error) {
		return roles.GetRolesServiceID(id)
	}
}

func RolesPostService(data interface{}) error {
	return roles.PostRolesService(data)
}

func RolesPutService(data interface{}) error {
	return roles.PutRolesService(data)
}

func RolesDeleteService(id int) error {
	return roles.DeleteRolesService(id)
}
