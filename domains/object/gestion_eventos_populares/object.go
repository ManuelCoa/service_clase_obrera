package eventos_populares

import "time"

type GestionEventosPopulares struct {
	IDEvento          int       `json:"id_evento"`
	IDGeoEventos      int       `json:"id_geo_eventos"`
	IDTipoEventos     int       `json:"id_tipo_eventos"`
	IDInstitucion     int       `json:"id_institucion"`
	NombreEvento      string    `json:"nombre_evento"`
	CuotaPersonas     string    `json:"cuota_personas"`
	FechaEvento       time.Time `json:"fecha_evento"`
	HoraEvento        string    `json:"hora_evento"`
	DescripcionEvento string    `json:"descripcion_evento"`
	EventoActivo      int       `json:"evento_activo"`
}