package certificacion_extracademica

import (
	"time"
)

type CertificacionExtracademica struct {
	CertificacionID     int       `json:"id_certificacion"`
	NumeroCertificacion string    `json:"numero_certificacion"`
	FechaEgreso         time.Time `json:"fecha_egreso"`
}
