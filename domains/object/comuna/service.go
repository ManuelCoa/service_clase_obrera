package comuna

func GetComunaService() ([]Comuna, error) {
	return searchParsedComuna()
}

func GetComunaServiceID(id int) ([]Comuna, error) {
	return searchParsedComunaID(id)
}

func PostComunaService(data interface{}) error {
	return insertionComuna(data)
}

func PutComunaService(data interface{}) error {
	return putComuna(data)
}

func DeleteComunaService(id int) error {
	return deletionComuna(id)
}