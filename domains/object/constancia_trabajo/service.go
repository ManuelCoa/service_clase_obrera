package constancia_trabajo


func GetConstanciasService() ([]ConstanciaTrabajo, error) {
	return searchParsedConstancias()
}

func GetConstanciaServiceID(id int) ([]ConstanciaTrabajo, error) {
	return searchParsedConstanciaID(id)
}

func PostConstanciaService(data interface{}) error {
	return insertionConstancia(data)
}

func PutConstanciaService(data interface{}) error {
	return putConstancia(data)
}

func DeleteConstanciaService(id int) error {
	return deletionConstancia(id)
}