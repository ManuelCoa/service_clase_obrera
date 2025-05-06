package interfaz

import (
	"claseobrera/domains/object/asistencia_biometrica"
)

type AsistenciaBiometricaInterface func() ([]asistencia_biometrica.AsistenciaBiometrica, error)

func AsistenciaBiometricaGetService() AsistenciaBiometricaInterface {
	return func() ([]asistencia_biometrica.AsistenciaBiometrica, error) {
		return asistencia_biometrica.GetAsistenciasBiometricasService()
	}
}

func AsistenciaBiometricaGetServiceID(id int) AsistenciaBiometricaInterface {
	return func() ([]asistencia_biometrica.AsistenciaBiometrica, error) {
		return asistencia_biometrica.GetAsistenciaBiometricaServiceID(id)
	}
}

func AsistenciaBiometricaPostService(data interface{}) error {
	return asistencia_biometrica.PostAsistenciaBiometricaService(data)
}

func AsistenciaBiometricaPutService(data interface{}) error {
	return asistencia_biometrica.PutAsistenciaBiometricaService(data)
}

func AsistenciaBiometricaDeleteService(id int) error {
	return asistencia_biometrica.DeleteAsistenciaBiometricaService(id)
}