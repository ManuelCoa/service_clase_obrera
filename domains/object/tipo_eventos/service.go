package tipo_eventos

func GetTiposEventoService() ([]TipoEventos, error) {
    return searchParsedTiposEvento()
}

func GetTipoEventoServiceID(id int) ([]TipoEventos, error) {
    return searchParsedTipoEventoID(id)
}

func PostTipoEventoService(data interface{}) error {
    return insertionTipoEvento(data)
}

func PutTipoEventoService(data interface{}) error {
    return putTipoEvento(data)
}

func DeleteTipoEventoService(id int) error {
    return deletionTipoEvento(id)
}

