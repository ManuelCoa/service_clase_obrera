package interfaz

import "claseobrera/domains/object/tipo_eventos"

type TipoEventoInterface func() ([]tipo_eventos.TipoEventos, error)

func TipoEventoGetService() TipoEventoInterface {
    return func() ([]tipo_eventos.TipoEventos, error) {
        return tipo_eventos.GetTiposEventoService()
    }
}

func TipoEventoGetServiceID(id int) TipoEventoInterface {
    return func() ([]tipo_eventos.TipoEventos, error) {
        return tipo_eventos.GetTipoEventoServiceID(id)
    }
}

func TipoEventoPostService(data interface{}) error {
    return tipo_eventos.PostTipoEventoService(data)
}

func TipoEventoPutService(data interface{}) error {
    return tipo_eventos.PutTipoEventoService(data)
}

func TipoEventoDeleteService(id int) error {
    return tipo_eventos.DeleteTipoEventoService(id)
}