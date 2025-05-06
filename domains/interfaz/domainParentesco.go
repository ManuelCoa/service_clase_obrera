package interfaz

import (
	"claseobrera/domains/object/parentesco"
)

type ParentescoInterface func() ([]parentesco.Parentesco, error)

func ParentescoGetService() ParentescoInterface {
	return func() ([]parentesco.Parentesco, error) {
		return parentesco.GetParentescoService()
	}
}

func ParentescoGetServiceID(id int) ParentescoInterface {
	return func() ([]parentesco.Parentesco, error) {
		return parentesco.GetParentescoServiceID(id)
	}
}

func ParentescoPostService(data interface{}) error {
	return parentesco.PostParentescoService(data)
}

func ParentescoPutService(data interface{}) error {
	return parentesco.PutParentescoService(data)
}

func ParentescoDeleteService(id int) error {
	return parentesco.DeleteParentescoService(id)
}