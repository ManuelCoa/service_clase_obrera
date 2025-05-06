package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetInstitucionesAcademicas() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.InstitucionAcademicaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetInstitucionAcademicaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.InstitucionAcademicaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostInstitucionAcademica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionAcademicaPostService(data)
	})
}

func PutInstitucionAcademica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionAcademicaPutService(data)
	})
}

func DeleteInstitucionAcademica(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionAcademicaDeleteService(id)
	})
}