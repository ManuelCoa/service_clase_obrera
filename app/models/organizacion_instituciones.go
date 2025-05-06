package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetInstituciones() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		org := interfaz.InstitucionesService()
		result, err := org()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetInstitucionesID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		orgID := interfaz.InstitucionesIDService(id)
		result, err := orgID()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostInstituciones(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionesPostService(data)
	})
}

func PutInstituciones(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionesPutService(data)
	})
}

func DeleteInstituciones(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.InstitucionesDeleteService(id)
	})
}
