package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetCentroVotacion() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CentroVotacionGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetCentroVotacionID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CentroVotacionGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostCentroVotacion(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CentroVotacionPostService(data)
	})
}

func PutCentroVotacion(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CentroVotacionPutService(data)
	})
}

func DeleteCentroVotacion(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CentroVotacionDeleteService(id)
	})
}
