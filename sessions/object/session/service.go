package session

func GetSessionGeneralService() ([]Session, error) {
	return searchParsedSession()
}

func GetSessionGeneralServiceID(id int) ([]Session, error) {
	return searchParsedSessionID(id)
}

func PostSessionGeneralService(data interface{}) error {
	return insertionSession(data)
}

func DeleteSessionGeneralService(token string) error {
	return deletionSession(token)
}
