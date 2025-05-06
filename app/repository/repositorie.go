package repository

func FetchDomainService(serviceFunc func() (interface{}, error)) (interface{}, error) {
	data, err := serviceFunc()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ModifyDomainService(serviceFunc func() error) error {
	err := serviceFunc()
	if err != nil {
		return err
	}
	return nil
}

func FetchSessionsService(serviceFunc func() (string, error)) (string, error) {
	data, err := serviceFunc()
	if err != nil {
		return "", err
	}
	return data, nil
}
