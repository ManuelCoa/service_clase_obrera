package tipo_parentesco

func GetTiposParentescoService() ([]TipoParentesco, error) {
	return searchParsedTiposParentesco()
}

func GetTiposParentescoServiceID(id int) ([]TipoParentesco, error) {
	return searchParsedTiposParentescoID(id)
}

func PostTipoParentescoService(data interface{}) error {
	return insertionTipoParentesco(data)
}

func PutTipoParentescoService(data interface{}) error {
	return putTipoParentesco(data)
}

func DeleteTipoParentescoService(id int) error {
	return deletionTipoParentesco(id)
}