package interfaz

import (
	gestionEventosPopulares "claseobrera/domains/object/gestion_eventos_populares"
)

type EventoInterface func() ([]gestionEventosPopulares.GestionEventosPopulares, error)

func EventoGetService() EventoInterface {
    return func() ([]gestionEventosPopulares.GestionEventosPopulares, error) {
        return gestionEventosPopulares.GetEventosService()
    }
}

func EventoGetServiceID(id int) EventoInterface {
    return func() ([]gestionEventosPopulares.GestionEventosPopulares, error) {
        return gestionEventosPopulares.GetEventoServiceID(id)
    }
}

func EventoPostService(data interface{}) error {
    return gestionEventosPopulares.PostEventoService(data)
}

func EventoPutService(data interface{}) error {
    return gestionEventosPopulares.PutEventoService(data)
}

func EventoDeleteService(id int) error {
    return gestionEventosPopulares.DeleteEventoService(id)
}