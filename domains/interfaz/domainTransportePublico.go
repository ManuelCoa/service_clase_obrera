package interfaz

import (
	transportepublico "claseobrera/domains/object/gestion_transporte_publico"
)

type TransportePublicoInterface func() ([]transportepublico.TransportePublico, error)

func TransportePublicoGetService() TransportePublicoInterface {
	return func() ([]transportepublico.TransportePublico, error) {
		return transportepublico.GetTransportePublicoService()
	}
}

func TransportePublicoGetServiceID(id int) TransportePublicoInterface {
	return func() ([]transportepublico.TransportePublico, error) {
		return transportepublico.GetTransportePublicoServiceID(id)
	}
}

func TransportePublicoPostService(data interface{}) error {
	return transportepublico.PostTransportePublicoService(data)
}

func TransportePublicoPutService(data interface{}) error {
	return transportepublico.PutTransportePublicoService(data)
}

func TransportePublicoDeleteService(id int) error {
	return transportepublico.DeleteTransportePublicoService(id)
}