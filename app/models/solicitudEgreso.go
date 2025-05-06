package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetSolicitudesEgreso() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.SolicitudEgresoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetSolicitudEgresoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.SolicitudEgresoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostSolicitudEgreso(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SolicitudEgresoPostService(data)
	})
}

func PutSolicitudEgreso(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SolicitudEgresoPutService(data)
	})
}

func DeleteSolicitudEgreso(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SolicitudEgresoDeleteService(id)
	})
}