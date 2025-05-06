package gestionextracademica

func GetGestionExtracademicaService() ([]GestionExtracademica, error) {
	return searchParsedGestionExtracademica()
}

func GetGestionExtracademicaServiceID(id int) ([]GestionExtracademica, error) {
	return searchParsedGestionExtracademicaID(id)
}

func PostGestionExtracademicaService(data interface{}) error {
	return insertionGestionExtracademica(data)
}

func PutGestionExtracademicaService(data interface{}) error {
	return putGestionExtracademica(data)
}

func DeleteGestionExtracademicaService(id int) error {
	return deletionGestionExtracademica(id)
}