package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetMunicipio() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.MunicipioGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetMunicipioID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.MunicipioGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostMunicipio(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MunicipioPostService(data)
	})
}

func PutMunicipio(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MunicipioPutService(data)
	})
}

func DeleteMunicipio(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.MunicipioDeleteService(id)
	})
}
