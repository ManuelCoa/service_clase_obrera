package marca_vehiculo

func GetMarcasService() ([]MarcaVehiculo, error) {
	return searchParsedMarcas()
}

func GetMarcasServiceID(id int) ([]MarcaVehiculo, error) {
	return searchParsedMarcasID(id)
}

func PostMarcaService(data interface{}) error {
	return insertionMarca(data)
}

func PutMarcaService(data interface{}) error {
	return putMarca(data)
}

func DeleteMarcaService(id int) error {
	return deletionMarca(id)
}