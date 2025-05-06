package interfaz

import (
	"claseobrera/sessions/object/perfiles"
)

type PerfilesInterface func() ([]perfiles.Perfiles, error)

func PerfilesGetService() PerfilesInterface {
	return func() ([]perfiles.Perfiles, error) {
		return perfiles.GetPerfilesService()
	}
}

func PerfilesPostService(data interface{}) error {
	return perfiles.PostPerfilesService(data)
}

func PerfilesDeleteService(rolesID, permisosID int) error {
	return perfiles.DeletePerfilesService(rolesID, permisosID)
}
