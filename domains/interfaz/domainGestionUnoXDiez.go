package interfaz

import (
	"claseobrera/domains/object/gestion_unoxdiez"
)

type GestionUnoxDiezInterface func() ([]gestion_unoxdiez.GestionUnoxDiez, error)

func GestionUnoxDiezGetService() GestionUnoxDiezInterface {
	return func() ([]gestion_unoxdiez.GestionUnoxDiez, error) {
		return gestion_unoxdiez.GetGestionUnoxDiezService()
	}
}

func GestionUnoxDiezGetServiceID(id int) GestionUnoxDiezInterface {
	return func() ([]gestion_unoxdiez.GestionUnoxDiez, error) {
		return gestion_unoxdiez.GetGestionUnoxDiezServiceID(id)
	}
}

func GestionUnoxDiezPostService(data interface{}) error {
	return gestion_unoxdiez.PostGestionUnoxDiezService(data)
}

func GestionUnoxDiezPutService(data interface{}) error {
	return gestion_unoxdiez.PutGestionUnoxDiezService(data)
}

func GestionUnoxDiezDeleteService(id int) error {
	return gestion_unoxdiez.DeleteGestionUnoxDiezService(id)
}