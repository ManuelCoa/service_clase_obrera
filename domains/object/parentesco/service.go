package parentesco

func GetParentescoService() ([]Parentesco, error) {
	return searchParsedParentesco()
}

func GetParentescoServiceID(id int) ([]Parentesco, error) {
	return searchParsedParentescoID(id)
}

func PostParentescoService(data interface{}) error {
	return insertionParentesco(data)
}

func PutParentescoService(data interface{}) error {
	return putParentesco(data)
}

func DeleteParentescoService(id int) error {
	return deletionParentesco(id)
}