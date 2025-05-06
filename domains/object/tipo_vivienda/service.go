package tipo_vivienda

func GetTiposViviendaService() ([]TipoVivienda, error) {
	return searchParsedTiposVivienda()
}

func GetTiposViviendaServiceID(id int) ([]TipoVivienda, error) {
	return searchParsedTiposViviendaID(id)
}

func PostTipoViviendaService(data interface{}) error {
	return insertionTipoVivienda(data)
}

func PutTipoViviendaService(data interface{}) error {
	return putTipoVivienda(data)
}

func DeleteTipoViviendaService(id int) error {
	return deletionTipoVivienda(id)
}