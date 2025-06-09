package test_piel

func GetTestPielService() ([]TestPiel, error) {
	return searchParsedTestPiel()
}

func GetTestPielServiceID(id int) ([]TestPiel, error) {
	return searchParsedTestPielID(id)
}

func PostTestPielService(data interface{}) error {
	return insertionTestPiel(data)
}

func PutTestPielService(data interface{}) error {
	return putTestPiel(data)
}

func DeleteTestPielService(id int) error {
	return deletionTestPiel(id)
}
