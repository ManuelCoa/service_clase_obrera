package estados

func GetEstadosService() ([]Estado, error) {
	return searchParsedEstado()
}

func GetEstadoIDService(id int) ([]Estado, error) {
	return searchParsedEstadoID(id)
}

func PostEstadoService(data interface{}) error {
	return insertionEstado(data)
}

func PutEstadoService(data interface{}) error {
	return putEstado(data)
}

func DeleteEstadoService(id int) error {
	return deletionEstado(id)
}
