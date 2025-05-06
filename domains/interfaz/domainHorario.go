package interfaz

import (
	"claseobrera/domains/object/horario"
)

type HorarioInterface func() ([]horario.Horario, error)

func HorarioGetService() HorarioInterface {
	return func() ([]horario.Horario, error) {
		return horario.GetHorariosService()
	}
}

func HorarioGetServiceID(id int) HorarioInterface {
	return func() ([]horario.Horario, error) {
		return horario.GetHorariosServiceID(id)
	}
}

func HorarioPostService(data interface{}) error {
	return horario.PostHorarioService(data)
}

func HorarioPutService(data interface{}) error {
	return horario.PutHorarioService(data)
}

func HorarioDeleteService(id int) error {
	return horario.DeleteHorarioService(id)
}