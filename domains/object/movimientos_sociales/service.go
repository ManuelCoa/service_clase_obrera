package movimientos_sociales

func GetMovimientosSocialesService() ([]MovimientosSociales, error) {
	return searchParsedMovimientosSociales()
}

func GetMovimientosSocialesServiceID(id int) ([]MovimientosSociales, error) {
	return searchParsedMovimientosSocialesID(id)
}

func PostMovimientosSocialesService(data interface{}) error {
	return insertionMovimientosSociales(data)
}

func PutMovimientosSocialesService(data interface{}) error {
	return putMovimientosSociales(data)
}

func DeleteMovimientosSocialesService(id int) error {
	return deletionMovimientosSociales(id)
}