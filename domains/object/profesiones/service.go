package profesiones

func GetProfesionesService() ([]Profesiones, error) {
	return searchParsedProfesiones()
}

func GetProfesionesServiceID(id int) ([]Profesiones, error) {
	return searchParsedProfesionesID(id)
}

func PostProfesionesService(data interface{}) error {
	return insertionProfesiones(data)
}

func PutProfesionesService(data interface{}) error {
	return putProfesiones(data)
}

func DeleteProfesionesService(id int) error {
	return deletionProfesiones(id)
}