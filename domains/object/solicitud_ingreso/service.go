package solicitud_ingreso

func GetSolicitudIngresoService() ([]SolicitudIngreso, error) {
	return searchParsedSolicitudIngreso()
}

func GetSolicitudIngresoServiceID(id int) ([]SolicitudIngreso, error) {
	return searchParsedSolicitudIngresoID(id)
}

func PostSolicitudIngresoService(data interface{}) error {
	return insertionSolicitudIngreso(data)
}

func PutSolicitudIngresoService(data interface{}) error {
	return putSolicitudIngreso(data)
}

func DeleteSolicitudIngresoService(id int) error {
	return deletionSolicitudIngreso(id)
}