package gestion_parentesco

func GetParentescoService() ([]GestionParentesco, error) {
	return searchParsedParentesco()
}

func GetParentescoServiceID(id int) ([]GestionParentesco, error) {
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
