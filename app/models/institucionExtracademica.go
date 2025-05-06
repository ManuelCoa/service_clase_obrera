package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetInstitucionExtracademica() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.InstitucionExtracademicaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetInstitucionExtracademicaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.InstitucionExtracademicaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostInstitucionExtracademica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionExtracademicaPostService(data)
	})
}

func PutInstitucionExtracademica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionExtracademicaPutService(data)
	})
}

func DeleteInstitucionExtracademica(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionExtracademicaDeleteService(id)
	})
}