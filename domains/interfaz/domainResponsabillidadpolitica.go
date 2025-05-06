package interfaz

import (
	"claseobrera/domains/object/responsabilidad_politica"
)

type ResponsabilidadPoliticaInterface func() ([]responsabilidad_politica.ResponsabilidadPolitica, error)

func ResponsabilidadPoliticaGetService() ResponsabilidadPoliticaInterface {
	return func() ([]responsabilidad_politica.ResponsabilidadPolitica, error) {
		return responsabilidad_politica.GetResponsabilidadPoliticaService()
	}
}

func ResponsabilidadPoliticaGetServiceID(id int) ResponsabilidadPoliticaInterface {
	return func() ([]responsabilidad_politica.ResponsabilidadPolitica, error) {
		return responsabilidad_politica.GetResponsabilidadPoliticaServiceID(id)
	}
}

func ResponsabilidadPoliticaPostService(data interface{}) error {
	return responsabilidad_politica.PostResponsabilidadPoliticaService(data)
}

func ResponsabilidadPoliticaPutService(data interface{}) error {
	return responsabilidad_politica.PutResponsabilidadPoliticaService(data)
}

func ResponsabilidadPoliticaDeleteService(id int) error {
	return responsabilidad_politica.DeleteResponsabilidadPoliticaService(id)
}
