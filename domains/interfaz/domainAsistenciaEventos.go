package interfaz

import asistencia_eventos "claseobrera/domains/object/gestion_asistencia_eventos"

type AsistenciaEventoInterface func() ([]asistencia_eventos.GestionAsistenciaEventos, error)

func AsistenciaEventoGetService() AsistenciaEventoInterface {
    return func() ([]asistencia_eventos.GestionAsistenciaEventos, error) {
        return asistencia_eventos.GetAsistenciasService()
    }
}

func AsistenciaEventoGetServiceID(id int) AsistenciaEventoInterface {
    return func() ([]asistencia_eventos.GestionAsistenciaEventos, error) {
        return asistencia_eventos.GetAsistenciaServiceID(id)
    }
}

func AsistenciaEventoPostService(data interface{}) error {
    return asistencia_eventos.PostAsistenciaService(data)
}

func AsistenciaEventoPutService(data interface{}) error {
    return asistencia_eventos.PutAsistenciaService(data)
}

func AsistenciaEventoDeleteService(id int) error {
    return asistencia_eventos.DeleteAsistenciaService(id)
}