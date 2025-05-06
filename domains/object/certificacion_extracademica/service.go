package certificacion_extracademica

func GetCertificacionesService() ([]CertificacionExtracademica, error) {
	return searchParsedCertificaciones()
}

func GetCertificacionServiceID(id int) ([]CertificacionExtracademica, error) {
	return searchParsedCertificacionID(id)
}

func PostCertificacionService(data interface{}) error {
	return insertionCertificacion(data)
}

func PutCertificacionService(data interface{}) error {
	return putCertificacion(data)
}

func DeleteCertificacionService(id int) error {
	return deletionCertificacion(id)
}
