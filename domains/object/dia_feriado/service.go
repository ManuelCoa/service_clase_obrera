package dia_feriado

func GetDiaFeriadoService() ([]DiaFeriado, error) {
	return searchParsedDiaFeriado()
}

func GetDiaFeriadoServiceID(id int) ([]DiaFeriado, error) {
	return searchParsedDiaFeriadoID(id)
}

func PostDiaFeriadoService(data interface{}) error {
	return insertionDiaFeriado(data)
}

func PutDiaFeriadoService(data interface{}) error {
	return putDiaFeriado(data)
}

func DeleteDiaFeriadoService(id int) error {
	return deletionDiaFeriado(id)
}