package interfaz

import (
	"claseobrera/domains/object/departamentos"
)

type DepartamentosInterface func() ([]departamentos.Departamentos, error)

func DepartamentosGetService() DepartamentosInterface {
	return func() ([]departamentos.Departamentos, error) {
		return departamentos.GetDepartamentosService()
	}
}

func DepartamentosGetServiceID(id int) DepartamentosInterface {
	return func() ([]departamentos.Departamentos, error) {
		return departamentos.GetDepartamentosServiceID(id)
	}
}

func DepartamentosPostService(data interface{}) error {
	return departamentos.PostDepartamentosService(data)
}

func DepartamentosPutService(data interface{}) error {
	return departamentos.PutDepartamentosService(data)
}

func DepartamentosDeleteService(id int) error {
	return departamentos.DeleteDepartamentosService(id)
}