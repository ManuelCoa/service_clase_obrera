package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetBioquimicosObreros() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.BioquimicosObreroGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetBioquimicosObreroID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.BioquimicosObreroGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostBioquimicosObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.BioquimicosObreroPostService(data)
	})
}

func PutBioquimicosObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.BioquimicosObreroPutService(data)
	})
}

func DeleteBioquimicosObrero(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.BioquimicosObreroDeleteService(id)
	})
}