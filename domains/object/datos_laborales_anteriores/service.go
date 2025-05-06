package datos_laborales

func GetDatosLaboralesService() ([]DatosLaboralesAnteriores, error) {
	return searchParsedDatosLaborales()
}

func GetDatosLaboralesServiceID(id int) ([]DatosLaboralesAnteriores, error) {
	return searchParsedDatosLaboralesID(id)
}

func GetDatosLaboralesPorObreroService(cedula int) ([]DatosLaboralesAnteriores, error) {
	return searchParsedDatosPorObrero(cedula)
}

func PostDatosLaboralesService(data interface{}) error {
	return insertionDatosLaborales(data)
}

func PutDatosLaboralesService(data interface{}) error {
	return putDatosLaborales(data)
}

func DeleteDatosLaboralesService(id int) error {
	return deletionDatosLaborales(id)
}
