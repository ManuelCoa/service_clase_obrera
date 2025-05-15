package transporte_publico

func GetTransportePublicoService() ([]TransportePublico, error) {
    return searchParsedTransportePublico()
}

func GetTransportePublicoServiceID(id int) ([]TransportePublico, error) {
    return searchParsedTransportePublicoID(id)
}

func PostTransportePublicoService(data interface{}) error {
    return insertionTransportePublico(data)
}

func PutTransportePublicoService(data interface{}) error {
    return putTransportePublico(data)
}

func DeleteTransportePublicoService(id int) error {
    return deletionTransportePublico(id)
}