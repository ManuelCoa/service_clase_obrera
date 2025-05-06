package interfaz

import (
	"claseobrera/domains/object/gestion_horarios"
)

type GestionHorarioInterface func() ([]gestion_horarios.GestionHorario, error)

func GestionHorarioGetService() GestionHorarioInterface {
	return func() ([]gestion_horarios.GestionHorario, error) {
		return gestion_horarios.GetGestionHorariosService()
	}
}

func GestionHorarioGetServiceID(id int) GestionHorarioInterface {
	return func() ([]gestion_horarios.GestionHorario, error) {
		return gestion_horarios.GetGestionHorarioServiceID(id)
	}
}

func GestionHorarioPostService(data interface{}) error {
	return gestion_horarios.PostGestionHorarioService(data)
}

func GestionHorarioPutService(data interface{}) error {
	return gestion_horarios.PutGestionHorarioService(data)
}

func GestionHorarioDeleteService(id int) error {
	return gestion_horarios.DeleteGestionHorarioService(id)
}