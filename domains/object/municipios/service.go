package municipios

func GetMunicipioService() ([]Municipio, error) {
	return searchParsedMunicipios()
}

func GetMunicipioServiceID(id int) ([]Municipio, error) {
	return searchParsedMunicipioID(id)
}

func PostMunicipioService(data interface{}) error {
	return insertionMunicipio(data)
}

func PutMunicipioService(data interface{}) error {
	return putMunicipio(data)
}

func DeleteMunicipioService(id int) error {
	return deletionMunicipio(id)
}
