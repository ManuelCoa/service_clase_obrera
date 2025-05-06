package tipo_enfermedad

func GetTiposEnfermedadService() ([]TipoEnfermedad, error) {
	return searchParsedTiposEnfermedad()
}

func GetTiposEnfermedadServiceID(id int) ([]TipoEnfermedad, error) {
	return searchParsedTiposEnfermedadID(id)
}

func PostTipoEnfermedadService(data interface{}) error {
	return insertionTipoEnfermedad(data)
}

func PutTipoEnfermedadService(data interface{}) error {
	return putTipoEnfermedad(data)
}

func DeleteTipoEnfermedadService(id int) error {
	return deletionTipoEnfermedad(id)
}