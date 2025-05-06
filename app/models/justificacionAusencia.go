package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetJustificacionesAusencia() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.JustificacionAusenciaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetJustificacionAusenciaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.JustificacionAusenciaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostJustificacionAusencia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.JustificacionAusenciaPostService(data)
	})
}

func PutJustificacionAusencia(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.JustificacionAusenciaPutService(data)
	})
}

func DeleteJustificacionAusencia(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.JustificacionAusenciaDeleteService(id)
	})
}