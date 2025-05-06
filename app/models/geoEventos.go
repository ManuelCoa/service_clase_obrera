package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetGeoEvento() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GeoEventoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetGeoEventoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.GeoEventoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostGeoEvento(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GeoEventoPostService(data)
	})
}

func PutGeoEvento(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GeoEventoPutService(data)
	})
}

func DeleteGeoEvento(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.GeoEventoDeleteService(id)
	})
}