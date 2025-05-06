package gestionextracademica

import (
	"time"
)
type GestionExtracademica struct {
	InfoExtracademicaID          int       `json:"id_info_extracademica"`
	ObreroCedula                 int       `json:"cedula_obrero"`
	TipoFormacionExtracademicaID int       `json:"id_tipo_formacion_extracademica"`
	CertificacionExtracademicaID int       `json:"id_certificacion_extracademica"`
	InstitucionExtracademicaID   int       `json:"id_institucion_extracademica"`
	FechaInicio                  time.Time `json:"fecha_inicio"`
	FechaFinalizacion            time.Time `json:"fecha_finalizacion"`
}