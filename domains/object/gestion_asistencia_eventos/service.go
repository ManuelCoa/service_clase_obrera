package asistencia_eventos

func GetAsistenciasService() ([]GestionAsistenciaEventos, error) {
    return searchParsedAsistencias()
}

func GetAsistenciaServiceID(id int) ([]GestionAsistenciaEventos, error) {
    return searchParsedAsistenciaID(id)
}

func PostAsistenciaService(data interface{}) error {
    return insertionAsistencia(data)
}

func PutAsistenciaService(data interface{}) error {
    return putAsistencia(data)
}

func DeleteAsistenciaService(id int) error {
    return deletionAsistencia(id)
}
