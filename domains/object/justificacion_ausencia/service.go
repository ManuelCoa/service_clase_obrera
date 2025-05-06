package justificacion_ausencia

func GetJustificacionesService() ([]JustificacionAusencia, error) {
	return searchParsedJustificaciones()
}

func GetJustificacionServiceID(id int) ([]JustificacionAusencia, error) {
	return searchParsedJustificacionID(id)
}

func PostJustificacionService(data interface{}) error {
	return insertionJustificacion(data)
}

func PutJustificacionService(data interface{}) error {
	return putJustificacion(data)
}

func DeleteJustificacionService(id int) error {
	return deletionJustificacion(id)
}
