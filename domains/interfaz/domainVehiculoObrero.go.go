package interfaz

import (
	"claseobrera/domains/object/vehiculo_obrero"
)

type VehiculoObreroInterface func() ([]vehiculo_obrero.VehiculoObrero, error)

func VehiculoObreroGetService() VehiculoObreroInterface {
	return func() ([]vehiculo_obrero.VehiculoObrero, error) {
		return vehiculo_obrero.GetVehiculoObreroService()
	}
}

func VehiculoObreroGetServiceID(id int) VehiculoObreroInterface {
	return func() ([]vehiculo_obrero.VehiculoObrero, error) {
		return vehiculo_obrero.GetVehiculoObreroServiceID(id)
	}
}

func VehiculoObreroPostService(data interface{}) error {
	return vehiculo_obrero.PostVehiculoObreroService(data)
}

func VehiculoObreroPutService(data interface{}) error {
	return vehiculo_obrero.PutVehiculoObreroService(data)
}

func VehiculoObreroDeleteService(id int) error {
	return vehiculo_obrero.DeleteVehiculoObreroService(id)
}