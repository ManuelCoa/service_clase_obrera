package amonestaciones

func GetAmonestacionesService() ([]Amonestaciones, error) {
	return searchParsedAmonestaciones()
}

func GetAmonestacionServiceID(id int) ([]Amonestaciones, error) {
	return searchParsedAmonestacionID(id)
}

func PostAmonestacionService(data interface{}) error {
	return insertionAmonestacion(data)
}

func PutAmonestacionService(data interface{}) error {
	return putAmonestacion(data)
}

func DeleteAmonestacionService(id int) error {
	return deletionAmonestacion(id)
}