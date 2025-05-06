package interfaz

import (
	"claseobrera/domains/object/consejo_comunal"
)

type ConsejoComunalInterface func() ([]consejo_comunal.ConsejoComunal, error)

func ConsejoComunalGetService() ConsejoComunalInterface {
	return func() ([]consejo_comunal.ConsejoComunal, error) {
		return consejo_comunal.GetConsejoComunalService()
	}
}

func ConsejoComunalGetServiceID(id int) ConsejoComunalInterface {
	return func() ([]consejo_comunal.ConsejoComunal, error) {
		return consejo_comunal.GetConsejoComunalServiceID(id)
	}
}

func ConsejoComunalPostService(data interface{}) error {
	return consejo_comunal.PostConsejoComunalService(data)
}

func ConsejoComunalPutService(data interface{}) error {
	return consejo_comunal.PutConsejoComunalService(data)
}

func ConsejoComunalDeleteService(id int) error {
	return consejo_comunal.DeleteConsejoComunalService(id)
}