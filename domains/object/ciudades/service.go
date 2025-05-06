package ciudades

func GetCiudadService() ([]Ciudad, error) {
	return searchParsedCiudad()
}

func GetCiudadServiceID(id int) ([]Ciudad, error) {
	return searchParsedCiudadID(id)
}

func PostCiudadService(data interface{}) error {
	return insertionCiudad(data)
}

func PutCiudadService(data interface{}) error {
	return putCiudad(data)
}

func DeleteCiudadService(id int) error {
	return deletionCiudad(id)
}
