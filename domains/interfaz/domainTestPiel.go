package interfaz

import (
	test_piel "claseobrera/domains/object/config_test_piel"
)

type TestPielInterface func() ([]test_piel.TestPiel, error)

func TestPielGetService() TestPielInterface {
	return func() ([]test_piel.TestPiel, error) {
		return test_piel.GetTestPielService()
	}
}

func TestPielGetServiceID(id int) TestPielInterface {
	return func() ([]test_piel.TestPiel, error) {
		return test_piel.GetTestPielServiceID(id)
	}
}

func TestPielPostService(data interface{}) error {
	return test_piel.PostTestPielService(data)
}

func TestPielPutService(data interface{}) error {
	return test_piel.PutTestPielService(data)
}

func TestPielDeleteService(id int) error {
	return test_piel.DeleteTestPielService(id)
}