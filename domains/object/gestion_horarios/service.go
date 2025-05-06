package gestion_horarios

func GetGestionHorariosService() ([]GestionHorario, error) {
	return searchParsedGestionHorarios()
}

func GetGestionHorarioServiceID(id int) ([]GestionHorario, error) {
	return searchParsedGestionHorarioID(id)
}

func PostGestionHorarioService(data interface{}) error {
	return insertionGestionHorario(data)
}

func PutGestionHorarioService(data interface{}) error {
	return putGestionHorario(data)
}

func DeleteGestionHorarioService(id int) error {
	return deletionGestionHorario(id)
}
