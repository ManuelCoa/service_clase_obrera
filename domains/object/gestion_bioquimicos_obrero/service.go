package bioquimicos_obrero

func GetBioquimicosObrerosService() ([]BioquimicosObrero, error) {
	return searchParsedBioquimicosObreros()
}

func GetBioquimicosObreroServiceID(id int) ([]BioquimicosObrero, error) {
	return searchParsedBioquimicosObreroID(id)
}

func PostBioquimicosObreroService(data interface{}) error {
	return insertionBioquimicosObrero(data)
}

func PutBioquimicosObreroService(data interface{}) error {
	return putBioquimicosObrero(data)
}

func DeleteBioquimicosObreroService(id int) error {
	return deletionBioquimicosObrero(id)
}
