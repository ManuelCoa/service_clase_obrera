package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetReposoMedico() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ReposoMedicoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetReposoMedicoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.ReposoMedicoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostReposoMedico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ReposoMedicoPostService(data)
	})
}

func PutReposoMedico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ReposoMedicoPutService(data)
	})
}

func DeleteReposoMedico(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.ReposoMedicoDeleteService(id)
	})
}