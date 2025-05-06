package organizacion_instituciones

func GetInstitucionesService() ([]Instituciones, error) {
	return searchParsedInstituciones()
}

func GetInstitucionesIDService(id int) ([]Instituciones, error) {
	return searchParsedInstitucionesID(id)
}

func PostInstitucionesService(data interface{}) error {
	return insertionInstituciones(data)
}

func PutCategoriaOrganizacionService(data interface{}) error {
	return putInstituciones(data)
}

func DeleteInstitucionesService(id int) error {
	return deletionInstituciones(id)
}
