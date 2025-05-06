package horario

func GetHorariosService() ([]Horario, error) {
	return searchParsedHorarios()
}

func GetHorariosServiceID(id int) ([]Horario, error) {
	return searchParsedHorariosID(id)
}

func PostHorarioService(data interface{}) error {
	return insertionHorario(data)
}

func PutHorarioService(data interface{}) error {
	return putHorario(data)
}

func DeleteHorarioService(id int) error {
	return deletionHorario(id)
}