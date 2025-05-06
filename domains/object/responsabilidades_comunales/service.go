package responsabilidades_comunales

func GetResponsabilidadesService() ([]ResponsabilidadComunal, error) {
	return searchParsedResponsabilidades()
}

func GetResponsabilidadesServiceID(id int) ([]ResponsabilidadComunal, error) {
	return searchParsedResponsabilidadesID(id)
}

func PostResponsabilidadService(data interface{}) error {
	return insertionResponsabilidad(data)
}

func PutResponsabilidadService(data interface{}) error {
	return putResponsabilidad(data)
}

func DeleteResponsabilidadService(id int) error {
	return deletionResponsabilidad(id)
}