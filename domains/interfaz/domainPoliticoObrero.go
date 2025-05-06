package interfaz

import (
	"claseobrera/domains/object/politico_obrero"
)

type PoliticoObreroInterface func() ([]politico_obrero.PoliticoObrero, error)

func PoliticoObreroGetService() PoliticoObreroInterface {
	return func() ([]politico_obrero.PoliticoObrero, error) {
		return politico_obrero.GetPoliticoObreroService()
	}
}

func PoliticoObreroGetServiceID(id int) PoliticoObreroInterface {
	return func() ([]politico_obrero.PoliticoObrero, error) {
		return politico_obrero.GetPoliticoObreroServiceID(id)
	}
}

func PoliticoObreroPostService(data interface{}) error {
	return politico_obrero.PostPoliticoObreroService(data)
}

func PoliticoObreroPutService(data interface{}) error {
	return politico_obrero.PutPoliticoObreroService(data)
}

func PoliticoObreroDeleteService(id int) error {
	return politico_obrero.DeletePoliticoObreroService(id)
}