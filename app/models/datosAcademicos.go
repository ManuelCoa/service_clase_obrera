package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetDatosAcademicos() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DatosAcademicosGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDatosAcademicosID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DatosAcademicosGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDatosAcademicosPorObrero(cedula int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DatosAcademicosPorObreroService(cedula)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostDatosAcademicos(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DatosAcademicosPostService(data)
	})
}

func PutDatosAcademicos(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DatosAcademicosPutService(data)
	})
}

func DeleteDatosAcademicos(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DatosAcademicosDeleteService(id)
	})
}