package parroquias

func GetParroquiaService() ([]Parroquia, error) {
	return searchParsedParroquia()
}

func GetParroquiaServiceID(id int) ([]Parroquia, error) {
	return searchParsedParroquiaID(id)
}

func PostParroquiaService(data interface{}) error {
	return insertionParroquia(data)
}

func PutParroquiaService(data interface{}) error {
	return putParroquia(data)
}

func DeleteParroquiaService(id int) error {
	return deletionParroquia(id)
}
