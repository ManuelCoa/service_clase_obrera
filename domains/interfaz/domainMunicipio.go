package interfaz

import (
	"claseobrera/domains/object/municipios"
)

type MunicipiosInterface func() ([]municipios.Municipio, error)

func MunicipioGetService() MunicipiosInterface {
	return func() ([]municipios.Municipio, error) {
		return municipios.GetMunicipioService()
	}
}

func MunicipioGetServiceID(id int) MunicipiosInterface {
	return func() ([]municipios.Municipio, error) {
		return municipios.GetMunicipioServiceID(id)
	}
}

func MunicipioPostService(data interface{}) error {
	return municipios.PostMunicipioService(data)
}

func MunicipioPutService(data interface{}) error {
	return municipios.PutMunicipioService(data)
}

func MunicipioDeleteService(id int) error {
	return municipios.DeleteMunicipioService(id)
}
