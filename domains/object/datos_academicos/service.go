package datos_academicos

func GetDatosAcademicosService() ([]DatosAcademicos, error) {
	return searchParsedDatosAcademicos()
}

func GetDatosAcademicosServiceID(id int) ([]DatosAcademicos, error) {
	return searchParsedDatosAcademicosID(id)
}

func GetDatosAcademicosPorObreroService(cedula int) ([]DatosAcademicos, error) {
	return searchParsedDatosPorObrero(cedula)
}

func PostDatosAcademicosService(data interface{}) error {
	return insertionDatosAcademicos(data)
}

func PutDatosAcademicosService(data interface{}) error {
	return putDatosAcademicos(data)
}

func DeleteDatosAcademicosService(id int) error {
	return deletionDatosAcademicos(id)
}
