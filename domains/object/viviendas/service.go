package viviendas

func GetViviendaService() ([]Vivienda, error) {
	return searchParsedVivienda()
}

func GetViviendaServiceID(id int) ([]Vivienda, error) {
	return searchParsedViviendaID(id)
}

func PostViviendaService(data interface{}) error {
	return insertionVivienda(data)
}

func PutViviendaService(data interface{}) error {
	return putVivienda(data)
}

func DeleteViviendaService(id int) error {
	return deletionVivienda(id)
}
