package interfaz

import (
	"claseobrera/domains/object/profesiones"
)

type ProfesionInterface func() ([]profesiones.Profesiones, error)

func ProfesionGetService() ProfesionInterface {
	return func() ([]profesiones.Profesiones, error) {
		return profesiones.GetProfesionesService()
	}
}

func ProfesionGetServiceID(id int) ProfesionInterface {
	return func() ([]profesiones.Profesiones, error) {
		return profesiones.GetProfesionesServiceID(id)
	}
}

func ProfesionPostService(data interface{}) error {
	return profesiones.PostProfesionesService(data)
}

func ProfesionPutService(data interface{}) error {
	return profesiones.PutProfesionesService(data)
}

func ProfesionDeleteService(id int) error {
	return profesiones.DeleteProfesionesService(id)
}