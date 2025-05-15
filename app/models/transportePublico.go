package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetTransportePublico() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TransportePublicoGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetTransportePublicoID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.TransportePublicoGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostTransportePublico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TransportePublicoPostService(data)
	})
}

func PutTransportePublico(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TransportePublicoPutService(data)
	})
}

func DeleteTransportePublico(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.TransportePublicoDeleteService(id)
	})
}