package tipo_techo

func GetTiposTechoService() ([]TipoTecho, error) {
	return searchParsedTiposTecho()
}

func GetTiposTechoServiceID(id int) ([]TipoTecho, error) {
	return searchParsedTiposTechoID(id)
}

func PostTipoTechoService(data interface{}) error {
	return insertionTipoTecho(data)
}

func PutTipoTechoService(data interface{}) error {
	return putTipoTecho(data)
}

func DeleteTipoTechoService(id int) error {
	return deletionTipoTecho(id)
}