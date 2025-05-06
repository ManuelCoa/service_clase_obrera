package estado_civil

func GetEstadoCivilService() ([]EstadoCivil, error) {
	return searchParsedEstadoCivil()
}

func GetEstadoCivilServiceID(id int) ([]EstadoCivil, error) {
	return searchParsedEstadoCivilID(id)
}

func PostEstadoCivilService(data interface{}) error {
	return insertionEstadoCivil(data)
}

func PutEstadoCivilService(data interface{}) error {
	return putEstadoCivil(data)
}

func DeleteEstadoCivilService(id int) error {
	return deletionEstadoCivil(id)
}