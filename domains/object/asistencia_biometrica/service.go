package asistencia_biometrica

func GetAsistenciasBiometricasService() ([]AsistenciaBiometrica, error) {
	return searchParsedAsistenciasBiometricas()
}

func GetAsistenciaBiometricaServiceID(id int) ([]AsistenciaBiometrica, error) {
	return searchParsedAsistenciaBiometricaID(id)
}

func PostAsistenciaBiometricaService(data interface{}) error {
	return insertionAsistenciaBiometrica(data)
}

func PutAsistenciaBiometricaService(data interface{}) error {
	return putAsistenciaBiometrica(data)
}

func DeleteAsistenciaBiometricaService(id int) error {
	return deletionAsistenciaBiometrica(id)
}
