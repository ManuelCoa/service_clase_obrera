package centro_votacion

func GetCentroVotacionService() ([]CentroVotacion, error) {
	return searchParsedCentroVotacion()
}

func GetCentroVotacionServiceID(id int) ([]CentroVotacion, error) {
	return searchParsedCentroVotacionID(id)
}

func PostCentroVotacionService(data interface{}) error {
	return insertionCentroVotacion(data)
}

func PutCentroVotacionService(data interface{}) error {
	return putCentroVotacion(data)
}

func DeleteCentroVotacionService(id int) error {
	return deletionCentroVotacion(id)
}