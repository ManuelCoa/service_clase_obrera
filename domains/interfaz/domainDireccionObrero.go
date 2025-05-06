package interfaz

import (
	direccionobrero "claseobrera/domains/object/direccion_obrero"
)

type DireccionObreroInterface func() ([]direccionobrero.DireccionObrero, error)

func DireccionObreroGetService() DireccionObreroInterface {
	return func() ([]direccionobrero.DireccionObrero, error) {
		return direccionobrero.GetDireccionObreroService()
	}
}

func DireccionObreroGetServiceID(id int) DireccionObreroInterface {
	return func() ([]direccionobrero.DireccionObrero, error) {
		return direccionobrero.GetDireccionObreroServiceID(id)
	}
}

func DireccionObreroPostService(data interface{}) error {
	return direccionobrero.PostDireccionObreroService(data)
}

func DireccionObreroPutService(data interface{}) error {
	return direccionobrero.PutDireccionObreroService(data)
}

func DireccionObreroDeleteService(id int) error {
	return direccionobrero.DeleteDireccionObreroService(id)
}