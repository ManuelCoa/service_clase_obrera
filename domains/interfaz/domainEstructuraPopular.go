package interfaz

import (
	"claseobrera/domains/object/estructura_popular"
)

type EstructuraPopularInterface func() ([]estructura_popular.EstructuraPopular, error)

func EstructuraPopularGetService() EstructuraPopularInterface {
	return func() ([]estructura_popular.EstructuraPopular, error) {
		return estructura_popular.GetEstructuraPopularService()
	}
}

func EstructuraPopularGetServiceID(id int) EstructuraPopularInterface {
	return func() ([]estructura_popular.EstructuraPopular, error) {
		return estructura_popular.GetEstructuraPopularServiceID(id)
	}
}

func EstructuraPopularPostService(data interface{}) error {
	return estructura_popular.PostEstructuraPopularService(data)
}

func EstructuraPopularPutService(data interface{}) error {
	return estructura_popular.PutEstructuraPopularService(data)
}

func EstructuraPopularDeleteService(id int) error {
	return estructura_popular.DeleteEstructuraPopularService(id)
}