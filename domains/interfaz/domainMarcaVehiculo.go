package interfaz

import (
	"claseobrera/domains/object/marca_vehiculo"
)

type MarcaVehiculoInterface func() ([]marca_vehiculo.MarcaVehiculo, error)

func MarcaVehiculoGetService() MarcaVehiculoInterface {
	return func() ([]marca_vehiculo.MarcaVehiculo, error) {
		return marca_vehiculo.GetMarcasService()
	}
}

func MarcaVehiculoGetServiceID(id int) MarcaVehiculoInterface {
	return func() ([]marca_vehiculo.MarcaVehiculo, error) {
		return marca_vehiculo.GetMarcasServiceID(id)
	}
}

func MarcaVehiculoPostService(data interface{}) error {
	return marca_vehiculo.PostMarcaService(data)
}

func MarcaVehiculoPutService(data interface{}) error {
	return marca_vehiculo.PutMarcaService(data)
}

func MarcaVehiculoDeleteService(id int) error {
	return marca_vehiculo.DeleteMarcaService(id)
}