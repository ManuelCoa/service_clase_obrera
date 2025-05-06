package datos_academicos

import (
	"time"
)

type DatosAcademicos struct {
	DatosAcademicosID      int       `json:"id_datos_academicos"`
	ObreroCedula           int       `json:"cedula_obrero"`
	NivelAcademicoID       int       `json:"id_nivel_academico"`
	InstitucionAcademicaID int       `json:"id_institucion_academica"`
	TituloAcademicoID      int       `json:"id_titulo_academico"`
	FechaInicio            time.Time `json:"fecha_inicio"`
	FechaFinalizacion      time.Time `json:"fecha_finalizacion"`
	CertificacionObtenida  string    `json:"certificacion_obtenida"`
}
