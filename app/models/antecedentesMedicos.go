package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetAntecedentesMedicos() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AntecedentesMedicosGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetAntecedenteMedicoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AntecedentesMedicosGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostAntecedenteMedico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AntecedentesMedicosPostService(data)
	})
}

func PutAntecedenteMedico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AntecedentesMedicosPutService(data)
	})
}

func DeleteAntecedenteMedico(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AntecedentesMedicosDeleteService(id)
	})
}