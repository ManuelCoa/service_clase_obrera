package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetAsistenciasBiometricas() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AsistenciaBiometricaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetAsistenciaBiometricaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.AsistenciaBiometricaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostAsistenciaBiometrica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaBiometricaPostService(data)
	})
}

func PutAsistenciaBiometrica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaBiometricaPutService(data)
	})
}

func DeleteAsistenciaBiometrica(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.AsistenciaBiometricaDeleteService(id)
	})
}