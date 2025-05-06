package tipo_formacion_extracademica

func GetTiposFormacionService() ([]TipoFormacionExtracademica, error) {
	return searchParsedTiposFormacion()
}

func GetTiposFormacionServiceID(id int) ([]TipoFormacionExtracademica, error) {
	return searchParsedTiposFormacionID(id)
}

func PostTipoFormacionService(data interface{}) error {
	return insertionTipoFormacion(data)
}

func PutTipoFormacionService(data interface{}) error {
	return putTipoFormacion(data)
}

func DeleteTipoFormacionService(id int) error {
	return deletionTipoFormacion(id)
}