package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTituloAcademico() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TituloAcademicoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTituloAcademicoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TituloAcademicoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTituloAcademico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TituloAcademicoPostService(data)
	})
}

func PutTituloAcademico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TituloAcademicoPutService(data)
	})
}

func DeleteTituloAcademico(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TituloAcademicoDeleteService(id)
	})
}