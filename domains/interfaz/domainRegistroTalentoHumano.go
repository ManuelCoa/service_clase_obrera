package interfaz

import (
	"claseobrera/domains/object/registro_talento_humano"
)

type RegistroTalentoHumanoInterface func() ([]registro_talento_humano.RegistroTalentoHumano, error)

func RegistroTalentoHumanoGetService() RegistroTalentoHumanoInterface {
	return func() ([]registro_talento_humano.RegistroTalentoHumano, error) {
		return registro_talento_humano.GetRegistroTalentoHumanoService()
	}
}

func RegistroTalentoHumanoGetServiceID(id int) RegistroTalentoHumanoInterface {
	return func() ([]registro_talento_humano.RegistroTalentoHumano, error) {
		return registro_talento_humano.GetRegistroTalentoHumanoServiceID(id)
	}
}

func RegistroTalentoHumanoPostService(data interface{}) error {
	return registro_talento_humano.PostRegistroTalentoHumanoService(data)
}

func RegistroTalentoHumanoPutService(data interface{}) error {
	return registro_talento_humano.PutRegistroTalentoHumanoService(data)
}

func RegistroTalentoHumanoDeleteService(id int) error {
	return registro_talento_humano.DeleteRegistroTalentoHumanoService(id)
}