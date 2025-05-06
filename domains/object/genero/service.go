package genero

func GetGeneroService() ([]Genero, error) {
	return searchParsedGenero()
}

func GetGeneroServiceID(id int) ([]Genero, error) {
	return searchParsedGeneroID(id)
}

func PostGeneroService(data interface{}) error {
	return insertionGenero(data)
}

func PutGeneroService(data interface{}) error {
	return putGenero(data)
}

func DeleteGeneroService(id int) error {
	return deletionGenero(id)
}