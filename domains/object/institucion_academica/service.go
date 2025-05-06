package institucion_academica

func GetInstitucionesAcademicasService() ([]InstitucionAcademica, error) {
	return searchParsedInstitucionesAcademicas()
}

func GetInstitucionAcademicaServiceID(id int) ([]InstitucionAcademica, error) {
	return searchParsedInstitucionAcademicaID(id)
}

func PostInstitucionAcademicaService(data interface{}) error {
	return insertionInstitucionAcademica(data)
}

func PutInstitucionAcademicaService(data interface{}) error {
	return putInstitucionAcademica(data)
}

func DeleteInstitucionAcademicaService(id int) error {
	return deletionInstitucionAcademica(id)
}
