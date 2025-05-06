package nivel_academico

func GetNivelAcademicoService() ([]NivelAcademico, error) {
	return searchParsedNivelAcademico()
}

func GetNivelAcademicoServiceID(id int) ([]NivelAcademico, error) {
	return searchParsedNivelAcademicoID(id)
}

func PostNivelAcademicoService(data interface{}) error {
	return insertionNivelAcademico(data)
}

func PutNivelAcademicoService(data interface{}) error {
	return putNivelAcademico(data)
}

func DeleteNivelAcademicoService(id int) error {
	return deletionNivelAcademico(id)
}