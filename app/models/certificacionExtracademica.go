package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetCertificaciones() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CertificacionGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetCertificacionID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.CertificacionGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostCertificacion(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CertificacionPostService(data)
	})
}

func PutCertificacion(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CertificacionPutService(data)
	})
}

func DeleteCertificacion(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.CertificacionDeleteService(id)
	})
}