package cargos_onapre

func GetCargoOnapreService() ([]CargosOnapre, error) {
	return searchParsedCargoOnapre()
}
func GetCargoOnapreServiceID(id int) ([]CargosOnapre, error) {
	return searchParsedCargoOnapreID(id)
}

func PostCargoOnapreService(data interface{}) error {
	return insertionCargoOnapre(data)
}

func PutCargoOnapreService(data interface{}) error {
	return putCargoOnapre(data)
}

func DeleteCargoOnapreService(id int) error {
	return deletionCargoOnapre(id)
}