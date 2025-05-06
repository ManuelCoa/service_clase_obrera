package gestion_unoxdiez

func GetGestionUnoxDiezService() ([]GestionUnoxDiez, error) {
	return searchParsedGestionUnoxDiez()
}

func GetGestionUnoxDiezServiceID(id int) ([]GestionUnoxDiez, error) {
	return searchParsedGestionUnoxDiezID(id)
}

func PostGestionUnoxDiezService(data interface{}) error {
	return insertionGestionUnoxDiez(data)
}

func PutGestionUnoxDiezService(data interface{}) error {
	return putGestionUnoxDiez(data)
}

func DeleteGestionUnoxDiezService(id int) error {
	return deletionGestionUnoxDiez(id)
}
