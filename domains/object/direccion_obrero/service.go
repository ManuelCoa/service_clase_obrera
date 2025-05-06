package direccion_obrero

func GetDireccionObreroService() ([]DireccionObrero, error) {
	return searchParsedDireccionObrero()
}

func GetDireccionObreroServiceID(id int) ([]DireccionObrero, error) {
	return searchParsedDireccionObreroID(id)
}

func PostDireccionObreroService(data interface{}) error {
	return insertionDireccionObrero(data)
}

func PutDireccionObreroService(data interface{}) error {
	return putDireccionObrero(data)
}

func DeleteDireccionObreroService(id int) error {
	return deletionDireccionObrero(id)
}