package geo_eventos

func GetGeoEventosService() ([]GestionGeoEventos, error) {
    return searchParsedGeoEventos()
}

func GetGeoEventoServiceID(id int) ([]GestionGeoEventos, error) {
    return searchParsedGeoEventoID(id)
}

func PostGeoEventoService(data interface{}) error {
    return insertionGeoEvento(data)
}

func PutGeoEventoService(data interface{}) error {
    return putGeoEvento(data)
}

func DeleteGeoEventoService(id int) error {
    return deletionGeoEvento(id)
}