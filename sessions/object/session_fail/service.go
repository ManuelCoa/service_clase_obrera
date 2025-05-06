package failsession

func GetFailedLoginService() ([]FailedLogin, error) {
	return searchFailedLogin()
}

func GetFailedLoginServiceID(id int) ([]FailedLogin, error) {
	return searchFailedLoginID(id)
}

func PostFailedLoginService(data interface{}) error {
	return insertionFailedLogin(data)
}
