package data_obrero

func GetObrerosService() ([]DataObrero, error) {
	return searchParsedObreros()
}

func GetObreroServiceID(cedula int) ([]DataObrero, error) {
	return searchParsedObreroID(cedula)
}

func GetObreroServiceIDInstitucion(id_institucion int) ([]DataObrero, error) {
	return searchParsedObreroIDInstitucion(id_institucion)
}

func PostObreroService(data interface{}) error {
	return insertionObrero(data)
}

func PutObreroService(data interface{}) error {
	return putObrero(data)
}

func DeleteObreroService(cedula int) error {
	return deletionObrero(cedula)
}