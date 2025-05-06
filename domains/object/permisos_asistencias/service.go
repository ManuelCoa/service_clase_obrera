package permisos_asistencias

func GetPermisoAsistenciaService() ([]PermisoAsistencia, error) {
	return searchParsedPermisoAsistencia()
}

func GetPermisoAsistenciaServiceID(id int) ([]PermisoAsistencia, error) {
	return searchParsedPermisoAsistenciaID(id)
}

func PostPermisoAsistenciaService(data interface{}) error {
	return insertionPermisoAsistencia(data)
}

func PutPermisoAsistenciaService(data interface{}) error {
	return putPermisoAsistencia(data)
}

func DeletePermisoAsistenciaService(id int) error {
	return deletionPermisoAsistencia(id)
}