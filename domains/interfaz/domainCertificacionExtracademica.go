package interfaz

import (
	certificaciones "claseobrera/domains/object/certificacion_extracademica"
)

type CertificacionInterface func() ([]certificaciones.CertificacionExtracademica, error)

func CertificacionGetService() CertificacionInterface {
	return func() ([]certificaciones.CertificacionExtracademica, error) {
		return certificaciones.GetCertificacionesService()
	}
}

func CertificacionGetServiceID(id int) CertificacionInterface {
	return func() ([]certificaciones.CertificacionExtracademica, error) {
		return certificaciones.GetCertificacionServiceID(id)
	}
}

func CertificacionPostService(data interface{}) error {
	return certificaciones.PostCertificacionService(data)
}

func CertificacionPutService(data interface{}) error {
	return certificaciones.PutCertificacionService(data)
}

func CertificacionDeleteService(id int) error {
	return certificaciones.DeleteCertificacionService(id)
}