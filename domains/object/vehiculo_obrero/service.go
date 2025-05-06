package vehiculo_obrero

func GetVehiculoObreroService() ([]VehiculoObrero, error) {
	return searchParsedVehiculoObrero()
}

func GetVehiculoObreroServiceID(id int) ([]VehiculoObrero, error) {
	return searchParsedVehiculoObreroID(id)
}

func PostVehiculoObreroService(data interface{}) error {
	return insertionVehiculoObrero(data)
}

func PutVehiculoObreroService(data interface{}) error {
	return putVehiculoObrero(data)
}

func DeleteVehiculoObreroService(id int) error {
	return deletionVehiculoObrero(id)
}
