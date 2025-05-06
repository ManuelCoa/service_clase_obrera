package eventos_populares

func GetEventosService() ([]GestionEventosPopulares, error) {
    return searchParsedEventos()
}

func GetEventoServiceID(id int) ([]GestionEventosPopulares, error) {
    return searchParsedEventoID(id)
}

func PostEventoService(data interface{}) error {
    return insertionEvento(data)
}

func PutEventoService(data interface{}) error {
    return putEvento(data)
}

func DeleteEventoService(id int) error {
    return deletionEvento(id)
}

