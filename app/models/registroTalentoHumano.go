package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetRegistrosTalentoHumano() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.RegistroTalentoHumanoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetRegistroTalentoHumanoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.RegistroTalentoHumanoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostRegistroTalentoHumano(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.RegistroTalentoHumanoPostService(data)
	})
}

func PutRegistroTalentoHumano(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.RegistroTalentoHumanoPutService(data)
	})
}

func DeleteRegistroTalentoHumano(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.RegistroTalentoHumanoDeleteService(id)
	})
}