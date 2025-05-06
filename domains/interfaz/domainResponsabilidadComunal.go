package interfaz

import (
	"claseobrera/domains/object/responsabilidades_comunales"
)

type ResponsabilidadComunalInterface func() ([]responsabilidades_comunales.ResponsabilidadComunal, error)

func ResponsabilidadComunalGetService() ResponsabilidadComunalInterface {
	return func() ([]responsabilidades_comunales.ResponsabilidadComunal, error) {
		return responsabilidades_comunales.GetResponsabilidadesService()
	}
}

func ResponsabilidadComunalGetServiceID(id int) ResponsabilidadComunalInterface {
	return func() ([]responsabilidades_comunales.ResponsabilidadComunal, error) {
		return responsabilidades_comunales.GetResponsabilidadesServiceID(id)
	}
}

func ResponsabilidadComunalPostService(data interface{}) error {
	return responsabilidades_comunales.PostResponsabilidadService(data)
}

func ResponsabilidadComunalPutService(data interface{}) error {
	return responsabilidades_comunales.PutResponsabilidadService(data)
}

func ResponsabilidadComunalDeleteService(id int) error {
	return responsabilidades_comunales.DeleteResponsabilidadService(id)
}