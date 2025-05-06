package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTipoFormacionExtracademica() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoFormacionExtracademicaGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTipoFormacionExtracademicaID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TipoFormacionExtracademicaGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTipoFormacionExtracademica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoFormacionExtracademicaPostService(data)
	})
}

func PutTipoFormacionExtracademica(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoFormacionExtracademicaPutService(data)
	})
}

func DeleteTipoFormacionExtracademica(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TipoFormacionExtracademicaDeleteService(id)
	})
}