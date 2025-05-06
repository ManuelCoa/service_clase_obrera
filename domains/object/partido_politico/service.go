package partido_politico

func GetPartidoPoliticoService() ([]PartidoPolitico, error) {
	return searchParsedPartidoPolitico()
}

func GetPartidoPoliticoServiceID(id int) ([]PartidoPolitico, error) {
	return searchParsedPartidoPoliticoID(id)
}

func PostPartidoPoliticoService(data interface{}) error {
	return insertionPartidoPolitico(data)
}

func PutPartidoPoliticoService(data interface{}) error {
	return putPartidoPolitico(data)
}

func DeletePartidoPoliticoService(id int) error {
	return deletionPartidoPolitico(id)
}