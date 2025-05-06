package fuerza_politica

func GetFuerzasService() ([]FuerzaPolitica, error) {
	return searchParsedFuerzas()
}

func GetFuerzasServiceID(id int) ([]FuerzaPolitica, error) {
	return searchParsedFuerzasID(id)
}

func PostFuerzaService(data interface{}) error {
	return insertionFuerza(data)
}

func PutFuerzaService(data interface{}) error {
	return putFuerza(data)
}

func DeleteFuerzaService(id int) error {
	return deletionFuerza(id)
}