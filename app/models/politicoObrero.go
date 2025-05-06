package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetPoliticosObreros() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.PoliticoObreroGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetPoliticoObreroID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.PoliticoObreroGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostPoliticoObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PoliticoObreroPostService(data)
	})
}

func PutPoliticoObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PoliticoObreroPutService(data)
	})
}

func DeletePoliticoObrero(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.PoliticoObreroDeleteService(id)
	})
}