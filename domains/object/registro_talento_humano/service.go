package registro_talento_humano

func GetRegistroTalentoHumanoService() ([]RegistroTalentoHumano, error) {
	return searchParsedRegistroTalentoHumano()
}

func GetRegistroTalentoHumanoServiceID(id int) ([]RegistroTalentoHumano, error) {
	return searchParsedRegistroTalentoHumanoID(id)
}

func PostRegistroTalentoHumanoService(data interface{}) error {
	return insertionRegistroTalentoHumano(data)
}

func PutRegistroTalentoHumanoService(data interface{}) error {
	return putRegistroTalentoHumano(data)
}

func DeleteRegistroTalentoHumanoService(id int) error {
	return deletionRegistroTalentoHumano(id)
}
