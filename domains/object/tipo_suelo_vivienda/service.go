package tipo_suelo_vivienda

func GetTiposSueloService() ([]TipoSueloVivienda, error) {
	return searchParsedTiposSuelo()
}

func GetTiposSueloServiceID(id int) ([]TipoSueloVivienda, error) {
	return searchParsedTiposSueloID(id)
}

func PostTipoSueloService(data interface{}) error {
	return insertionTipoSuelo(data)
}

func PutTipoSueloService(data interface{}) error {
	return putTipoSuelo(data)
}

func DeleteTipoSueloService(id int) error {
	return deletionTipoSuelo(id)
}