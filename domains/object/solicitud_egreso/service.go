package solicitud_egreso

func GetSolicitudEgresoService() ([]SolicitudEgreso, error) {
	return searchParsedSolicitudEgreso()
}

func GetSolicitudEgresoServiceID(id int) ([]SolicitudEgreso, error) {
	return searchParsedSolicitudEgresoID(id)
}

func PostSolicitudEgresoService(data interface{}) error {
	return insertionSolicitudEgreso(data)
}

func PutSolicitudEgresoService(data interface{}) error {
	return putSolicitudEgreso(data)
}

func DeleteSolicitudEgresoService(id int) error {
	return deletionSolicitudEgreso(id)
}
