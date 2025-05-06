package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetDiasFeriados() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DiaFeriadoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDiaFeriadoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DiaFeriadoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostDiaFeriado(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DiaFeriadoPostService(data)
	})
}

func PutDiaFeriado(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DiaFeriadoPutService(data)
	})
}

func DeleteDiaFeriado(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DiaFeriadoDeleteService(id)
	})
}