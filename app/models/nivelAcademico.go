package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetNivelAcademico() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.NivelAcademicoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetNivelAcademicoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.NivelAcademicoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostNivelAcademico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.NivelAcademicoPostService(data)
	})
}

func PutNivelAcademico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.NivelAcademicoPutService(data)
	})
}

func DeleteNivelAcademico(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.NivelAcademicoDeleteService(id)
	})
}
