package asistencia_biometrica

import (
	"time"
)

type AsistenciaBiometrica struct {
	AsistenciaBiometricaID   int       `json:"id_asistencia_biometrica"`
	CedulaObrero             int       `json:"cedula_obrero"`
	IdentificacionBiometrica string    `json:"identificacion_biometrica"`
	FechaRegistro            time.Time `json:"fecha_registro"`
	Activa                   bool      `json:"activa"`
}
