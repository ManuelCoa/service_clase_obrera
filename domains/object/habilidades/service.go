package habilidades

func GetHabilidadesService() ([]Habilidades, error) {
	return searchParsedHabilidades()
}

func GetHabilidadServiceID(id int) ([]Habilidades, error) {
	return searchParsedHabilidadID(id)
}

func PostHabilidadService(data interface{}) error {
	return insertionHabilidad(data)
}

func PutHabilidadService(data interface{}) error {
	return putHabilidad(data)
}

func DeleteHabilidadService(id int) error {
	return deletionHabilidad(id)
}
