package interfaz

import (
	"claseobrera/sessions/object/users"
)

type UsersInterface func() ([]users.Usuarios, error)

func UsersService() UsersInterface {
	return func() ([]users.Usuarios, error) {
		return users.GetUsuariosService()
	}
}

func UsersIDService(id int) UsersInterface {
	return func() ([]users.Usuarios, error) {
		return users.GetUsuariosServiceID(id)
	}
}

func UsersPostService(data interface{}) error {
	return users.PostUsuariosService(data)
}

func UsersPutService(data interface{}) error {
	return users.PutUsuariosService(data)
}

func UsersDeleteService(id int) error {
	return users.DeleteUsuariosService(id)
}
