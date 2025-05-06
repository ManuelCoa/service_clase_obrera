package departamentos

func GetDepartamentosService() ([]Departamentos, error) {
	return searchParsedDepartamentos()
}

func GetDepartamentosServiceID(id int) ([]Departamentos, error) {
	return searchParsedDepartamentosID(id)
}

func PostDepartamentosService(data interface{}) error {
	return insertionDepartamentos(data)
}

func PutDepartamentosService(data interface{}) error {
	return putDepartamentos(data)
}

func DeleteDepartamentosService(id int) error {
	return deletionDepartamentos(id)
}