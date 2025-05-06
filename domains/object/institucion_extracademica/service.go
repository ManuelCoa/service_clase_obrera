package institucion_extracademica

func GetInstitucionesService() ([]InstitucionExtracademica, error) {
	return searchParsedInstituciones()
}

func GetInstitucionesServiceID(id int) ([]InstitucionExtracademica, error) {
	return searchParsedInstitucionesID(id)
}

func PostInstitucionService(data interface{}) error {
	return insertionInstitucion(data)
}

func PutInstitucionService(data interface{}) error {
	return putInstitucion(data)
}

func DeleteInstitucionService(id int) error {
	return deletionInstitucion(id)
}