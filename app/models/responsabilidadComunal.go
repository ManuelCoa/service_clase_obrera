package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetResponsabilidadComunal() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ResponsabilidadComunalGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetResponsabilidadComunalID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ResponsabilidadComunalGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostResponsabilidadComunal(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ResponsabilidadComunalPostService(data)
	})
}

func PutResponsabilidadComunal(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ResponsabilidadComunalPutService(data)
	})
}

func DeleteResponsabilidadComunal(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ResponsabilidadComunalDeleteService(id)
	})
}