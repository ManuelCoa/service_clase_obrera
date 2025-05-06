package interfaz

import (
	bioquimicos_obrero "claseobrera/domains/object/gestion_bioquimicos_obrero"
)

type BioquimicosObreroInterface func() ([]bioquimicos_obrero.BioquimicosObrero, error)

func BioquimicosObreroGetService() BioquimicosObreroInterface {
	return func() ([]bioquimicos_obrero.BioquimicosObrero, error) {
		return bioquimicos_obrero.GetBioquimicosObrerosService()
	}
}

func BioquimicosObreroGetServiceID(id int) BioquimicosObreroInterface {
	return func() ([]bioquimicos_obrero.BioquimicosObrero, error) {
		return bioquimicos_obrero.GetBioquimicosObreroServiceID(id)
	}
}

func BioquimicosObreroPostService(data interface{}) error {
	return bioquimicos_obrero.PostBioquimicosObreroService(data)
}

func BioquimicosObreroPutService(data interface{}) error {
	return bioquimicos_obrero.PutBioquimicosObreroService(data)
}

func BioquimicosObreroDeleteService(id int) error {
	return bioquimicos_obrero.DeleteBioquimicosObreroService(id)
}