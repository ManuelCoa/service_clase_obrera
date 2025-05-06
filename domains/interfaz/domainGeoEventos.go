package interfaz

import geo_eventos "claseobrera/domains/object/gestion_geo_eventos"

type GeoEventoInterface func() ([]geo_eventos.GestionGeoEventos, error)

func GeoEventoGetService() GeoEventoInterface {
    return func() ([]geo_eventos.GestionGeoEventos, error) {
        return geo_eventos.GetGeoEventosService()
    }
}

func GeoEventoGetServiceID(id int) GeoEventoInterface {
    return func() ([]geo_eventos.GestionGeoEventos, error) {
        return geo_eventos.GetGeoEventoServiceID(id)
    }
}

func GeoEventoPostService(data interface{}) error {
    return geo_eventos.PostGeoEventoService(data)
}

func GeoEventoPutService(data interface{}) error {
    return geo_eventos.PutGeoEventoService(data)
}

func GeoEventoDeleteService(id int) error {
    return geo_eventos.DeleteGeoEventoService(id)
}