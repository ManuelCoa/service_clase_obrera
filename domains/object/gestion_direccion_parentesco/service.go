package direccion_parentesco

func GetDireccionParentescoService() ([]DireccionParentesco, error) {
    return searchParsedDireccionParentesco()
}

func GetDireccionParentescoServiceID(id int) ([]DireccionParentesco, error) {
    return searchParsedDireccionParentescoID(id)
}

func PostDireccionParentescoService(data interface{}) error {
    return insertionDireccionParentesco(data)
}

func PutDireccionParentescoService(data interface{}) error {
    return putDireccionParentesco(data)
}

func DeleteDireccionParentescoService(id int) error {
    return deletionDireccionParentesco(id)
}