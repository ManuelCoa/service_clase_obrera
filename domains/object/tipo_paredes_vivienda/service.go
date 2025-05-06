package paredes_vivienda

func GetTiposParedService() ([]TipoParedesVivienda, error) {
    return searchParsedTiposPared()
}

func GetTipoParedServiceID(id int) ([]TipoParedesVivienda, error) {
    return searchParsedTipoParedID(id)
}

func PostTipoParedService(data interface{}) error {
    return insertionTipoPared(data)
}

func PutTipoParedService(data interface{}) error {
    return putTipoPared(data)
}

func DeleteTipoParedService(id int) error {
    return deletionTipoPared(id)
}
