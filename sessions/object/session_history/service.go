package historysession

func GetLoginHistoryService() ([]LoginHistory, error) {
	return searchLoginHistory()
}

func GetLoginHistoryServiceID(id int) ([]LoginHistory, error) {
	return searchLoginHistoryID(id)
}

func PostLoginHistoryService(data interface{}) error {
	return insertionLoginHistory(data)
}
