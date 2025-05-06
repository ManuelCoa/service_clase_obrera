package interfaz

import (
	"claseobrera/domains/object/centro_votacion"
)

type CentroVotacionInterface func() ([]centro_votacion.CentroVotacion, error)

func CentroVotacionGetService() CentroVotacionInterface {
	return func() ([]centro_votacion.CentroVotacion, error) {
		return centro_votacion.GetCentroVotacionService()
	}
}

func CentroVotacionGetServiceID(id int) CentroVotacionInterface {
	return func() ([]centro_votacion.CentroVotacion, error) {
		return centro_votacion.GetCentroVotacionServiceID(id)
	}
}

func CentroVotacionPostService(data interface{}) error {
	return centro_votacion.PostCentroVotacionService(data)
}

func CentroVotacionPutService(data interface{}) error {
	return centro_votacion.PutCentroVotacionService(data)
}

func CentroVotacionDeleteService(id int) error {
	return centro_votacion.DeleteCentroVotacionService(id)
}
