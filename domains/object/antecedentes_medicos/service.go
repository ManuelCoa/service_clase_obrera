package antecedentes_medicos

func GetAntecedentesMedicosService() ([]AntecedentesMedicos, error) {
	return searchParsedAntecedentesMedicos()
}

func GetAntecedenteMedicoServiceID(id int) ([]AntecedentesMedicos, error) {
	return searchParsedAntecedenteMedicoID(id)
}

func PostAntecedenteMedicoService(data interface{}) error {
	return insertionAntecedenteMedico(data)
}

func PutAntecedenteMedicoService(data interface{}) error {
	return putAntecedenteMedico(data)
}

func DeleteAntecedenteMedicoService(id int) error {
	return deletionAntecedenteMedico(id)
}
