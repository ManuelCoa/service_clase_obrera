package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetAcademicaParentesco() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AcademicaParentescoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetAcademicaParentescoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AcademicaParentescoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostAcademicaParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AcademicaParentescoPostService(data)
	})
}

func PutAcademicaParentesco(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AcademicaParentescoPutService(data)
	})
}

func DeleteAcademicaParentesco(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AcademicaParentescoDeleteService(id)
	})
}