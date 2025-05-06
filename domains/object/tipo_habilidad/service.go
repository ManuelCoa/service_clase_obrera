package tipo_habilidad

func GetTiposHabilidadService() ([]TipoHabilidad, error) {
	return searchParsedTiposHabilidad()
}

func GetTiposHabilidadServiceID(id int) ([]TipoHabilidad, error) {
	return searchParsedTiposHabilidadID(id)
}

func PostTipoHabilidadService(data interface{}) error {
	return insertionTipoHabilidad(data)
}

func PutTipoHabilidadService(data interface{}) error {
	return putTipoHabilidad(data)
}

func DeleteTipoHabilidadService(id int) error {
	return deletionTipoHabilidad(id)
}