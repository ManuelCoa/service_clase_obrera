package academica_parentesco

func GetAcademicasParentescoService() ([]AcademicaParentesco, error) {
	return searchParsedAcademicasParentesco()
}

func GetAcademicaParentescoServiceID(id int) ([]AcademicaParentesco, error) {
	return searchParsedAcademicaParentescoID(id)
}

func PostAcademicaParentescoService(data interface{}) error {
	return insertionAcademicaParentesco(data)
}

func PutAcademicaParentescoService(data interface{}) error {
	return putAcademicaParentesco(data)
}

func DeleteAcademicaParentescoService(id int) error {
	return deletionAcademicaParentesco(id)
}