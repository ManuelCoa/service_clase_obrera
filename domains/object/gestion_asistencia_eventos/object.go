package asistencia_eventos

type GestionAsistenciaEventos struct {
	IDAsistenciaEventos int    `json:"id_asistencia_eventos"`
	IDEvento            int    `json:"id_evento"`
	CedulaObrero        int    `json:"cedula_obrero"`
	IDInstitucion       int    `json:"id_institucion"`
	Participacion       int    `json:"participacion"`
	Observacion         string `json:"observacion"`
}