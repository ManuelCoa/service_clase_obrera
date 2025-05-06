package interfaz

import (
	"claseobrera/domains/object/data_obrero"
)

type DataObreroInterface func() ([]data_obrero.DataObrero, error)

func DataObreroGetService() DataObreroInterface {
	return func() ([]data_obrero.DataObrero, error) {
		return data_obrero.GetObrerosService()
	}
}

func DataObreroGetServiceID(cedula int) DataObreroInterface {
	return func() ([]data_obrero.DataObrero, error) {
		return data_obrero.GetObreroServiceID(cedula)
	}
}

func DataObreroPostService(data interface{}) error {
	return data_obrero.PostObreroService(data)
}

func DataObreroPutService(data interface{}) error {
	return data_obrero.PutObreroService(data)
}

func DataObreroDeleteService(cedula int) error {
	return data_obrero.DeleteObreroService(cedula)
}