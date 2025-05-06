package asistencia

import "time"

type Asistencia struct {
	AsistenciaID           int       `json:"id_asistencia"`
	AsistenciaBiometricaID int       `json:"id_asistencia_biometrica"`
	FeriadoID              int       `json:"id_feriado"`
	Fecha                  time.Time `json:"fecha"`
	HoraEntrada            time.Time `json:"hora_entrada"`
	HoraSalida             time.Time `json:"hora_salida"`
}
