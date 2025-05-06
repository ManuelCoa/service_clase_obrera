package titulo_academico

func GetTituloAcademicoService() ([]TituloAcademico, error) {
	return searchParsedTituloAcademico()
}

func GetTituloAcademicoServiceID(id int) ([]TituloAcademico, error) {
	return searchParsedTituloAcademicoID(id)
}

func PostTituloAcademicoService(data interface{}) error {
	return insertionTituloAcademico(data)
}

func PutTituloAcademicoService(data interface{}) error {
	return putTituloAcademico(data)
}

func DeleteTituloAcademicoService(id int) error {
	return deletionTituloAcademico(id)
}
