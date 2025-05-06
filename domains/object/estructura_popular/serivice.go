package estructura_popular

func GetEstructuraPopularService() ([]EstructuraPopular, error) {
	return searchParsedEstructuraPopular()
}

func GetEstructuraPopularServiceID(id int) ([]EstructuraPopular, error) {
	return searchParsedEstructuraPopularID(id)
}

func PostEstructuraPopularService(data interface{}) error {
	return insertionEstructuraPopular(data)
}

func PutEstructuraPopularService(data interface{}) error {
	return putEstructuraPopular(data)
}

func DeleteEstructuraPopularService(id int) error {
	return deletionEstructuraPopular(id)
}