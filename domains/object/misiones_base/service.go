package misiones_base

func GetMisionesService() ([]MisionesBase, error) {
	return searchParsedMisiones()
}

func GetMisionesServiceID(id int) ([]MisionesBase, error) {
	return searchParsedMisionesID(id)
}

func PostMisionService(data interface{}) error {
	return insertionMision(data)
}

func PutMisionService(data interface{}) error {
	return putMision(data)
}

func DeleteMisionService(id int) error {
	return deletionMision(id)
}