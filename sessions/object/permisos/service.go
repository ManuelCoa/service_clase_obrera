package permisos

func GetPermisosService() ([]Permisos, error) {
	return searchParsedPermisos()
}

func GetPermisosIDService(id int) ([]Permisos, error) {
	return searchParsedPermisosID(id)
}

func PostPermisosService(data interface{}) error {
	return insertionPermisos(data)
}

func PutPermisosService(data interface{}) error {
	return putPermisos(data)
}

func DeletePermisosService(id int) error {
	return deletionPermisos(id)
}
