package cargo_publico

func GetCargoService() ([]CargoPublico, error) {
	return searchParsedCargo()
}

func GetCargoServiceID(id int) ([]CargoPublico, error) {
	return searchParsedCargoID(id)
}

func PostCargoService(data interface{}) error {
	return insertionCargo(data)
}

func PutCargoService(data interface{}) error {
	return putCargo(data)
}

func DeleteCargoService(id int) error {
	return deletionCargo(id)
}