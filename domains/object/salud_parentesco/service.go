package salud_parentesco

func GetSaludParentescoService() ([]SaludParentesco, error) {
	return searchParsedSaludParentesco()
}

func GetSaludParentescoServiceID(id int) ([]SaludParentesco, error) {
	return searchParsedSaludParentescoID(id)
}

func PostSaludParentescoService(data interface{}) error {
	return insertionSaludParentesco(data)
}

func PutSaludParentescoService(data interface{}) error {
	return putSaludParentesco(data)
}

func DeleteSaludParentescoService(id int) error {
	return deletionSaludParentesco(id)
}
