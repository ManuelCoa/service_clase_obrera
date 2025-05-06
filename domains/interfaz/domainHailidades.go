package interfaz

import (
	"claseobrera/domains/object/habilidades"
)

type HabilidadesInterface func() ([]habilidades.Habilidades, error)

func HabilidadesGetService() HabilidadesInterface {
	return func() ([]habilidades.Habilidades, error) {
		return habilidades.GetHabilidadesService()
	}
}

func HabilidadesGetServiceID(id int) HabilidadesInterface {
	return func() ([]habilidades.Habilidades, error) {
		return habilidades.GetHabilidadServiceID(id)
	}
}

func HabilidadesPostService(data interface{}) error {
	return habilidades.PostHabilidadService(data)
}

func HabilidadesPutService(data interface{}) error {
	return habilidades.PutHabilidadService(data)
}

func HabilidadesDeleteService(id int) error {
	return habilidades.DeleteHabilidadService(id)
}