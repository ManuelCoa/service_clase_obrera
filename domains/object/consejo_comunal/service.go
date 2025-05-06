package consejo_comunal

func GetConsejoComunalService() ([]ConsejoComunal, error) {
	return searchParsedConsejoComunal()
}

func GetConsejoComunalServiceID(id int) ([]ConsejoComunal, error) {
	return searchParsedConsejoComunalID(id)
}

func PostConsejoComunalService(data interface{}) error {
	return insertionConsejoComunal(data)
}

func PutConsejoComunalService(data interface{}) error {
	return putConsejoComunal(data)
}

func DeleteConsejoComunalService(id int) error {
	return deletionConsejoComunal(id)
}