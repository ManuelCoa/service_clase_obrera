package responsabilidad_politica

func GetResponsabilidadPoliticaService() ([]ResponsabilidadPolitica, error) {
	return searchParsedResponsabilidadPolitica()
}

func GetResponsabilidadPoliticaServiceID(id int) ([]ResponsabilidadPolitica, error) {
	return searchParsedResponsabilidadPoliticaID(id)
}

func PostResponsabilidadPoliticaService(data interface{}) error {
	return insertionResponsabilidadPolitica(data)
}

func PutResponsabilidadPoliticaService(data interface{}) error {
	return putResponsabilidadPolitica(data)
}

func DeleteResponsabilidadPoliticaService(id int) error {
	return deletionResponsabilidadPolitica(id)
}