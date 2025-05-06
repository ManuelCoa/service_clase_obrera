package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetSolicitudesIngreso() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.SolicitudIngresoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetSolicitudIngresoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.SolicitudIngresoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostSolicitudIngreso(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SolicitudIngresoPostService(data)
	})
}

func PutSolicitudIngreso(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SolicitudIngresoPutService(data)
	})
}

func DeleteSolicitudIngreso(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.SolicitudIngresoDeleteService(id)
	})
}