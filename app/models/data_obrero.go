package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetDataObrero() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DataObreroGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetDataObreroID(cedula int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DataObreroGetServiceID(cedula)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}
//fucnicon para consultar por id
func GetDataObreroIDInstitucion(id_institucion int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.DataObreroGetServiceIDInstitucion(id_institucion)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostDataObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DataObreroPostService(data)
	})
}

func PutDataObrero(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DataObreroPutService(data)
	})
}

func DeleteDataObrero(cedula int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.DataObreroDeleteService(cedula)
	})
}