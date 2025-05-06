package politico_obrero

func GetPoliticoObreroService() ([]PoliticoObrero, error) {
	return searchParsedPoliticoObrero()
}

func GetPoliticoObreroServiceID(id int) ([]PoliticoObrero, error) {
	return searchParsedPoliticoObreroID(id)
}

func PostPoliticoObreroService(data interface{}) error {
	return insertionPoliticoObrero(data)
}

func PutPoliticoObreroService(data interface{}) error {
	return putPoliticoObrero(data)
}

func DeletePoliticoObreroService(id int) error {
	return deletionPoliticoObrero(id)
}
