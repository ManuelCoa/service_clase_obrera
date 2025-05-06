package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetResponsabilidadPolitica() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ResponsabilidadPoliticaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetResponsabilidadPoliticaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ResponsabilidadPoliticaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostResponsabilidadPolitica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ResponsabilidadPoliticaPostService(data)
	})
}

func PutResponsabilidadPolitica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ResponsabilidadPoliticaPutService(data)
	})
}

func DeleteResponsabilidadPolitica(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ResponsabilidadPoliticaDeleteService(id)
	})
}
