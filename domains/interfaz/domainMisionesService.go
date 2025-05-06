package interfaz

import (
	"claseobrera/domains/object/misiones_base"
)

type MisionesBaseInterface func() ([]misiones_base.MisionesBase, error)

func MisionesBaseGetService() MisionesBaseInterface {
	return func() ([]misiones_base.MisionesBase, error) {
		return misiones_base.GetMisionesService()
	}
}

func MisionesBaseGetServiceID(id int) MisionesBaseInterface {
	return func() ([]misiones_base.MisionesBase, error) {
		return misiones_base.GetMisionesServiceID(id)
	}
}

func MisionesBasePostService(data interface{}) error {
	return misiones_base.PostMisionService(data)
}

func MisionesBasePutService(data interface{}) error {
	return misiones_base.PutMisionService(data)
}

func MisionesBaseDeleteService(id int) error {
	return misiones_base.DeleteMisionService(id)
}