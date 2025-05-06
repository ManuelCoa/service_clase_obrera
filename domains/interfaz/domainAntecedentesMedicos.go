package interfaz

import (
	"claseobrera/domains/object/antecedentes_medicos"
)

type AntecedentesMedicosInterface func() ([]antecedentes_medicos.AntecedentesMedicos, error)

func AntecedentesMedicosGetService() AntecedentesMedicosInterface {
	return func() ([]antecedentes_medicos.AntecedentesMedicos, error) {
		return antecedentes_medicos.GetAntecedentesMedicosService()
	}
}

func AntecedentesMedicosGetServiceID(id int) AntecedentesMedicosInterface {
	return func() ([]antecedentes_medicos.AntecedentesMedicos, error) {
		return antecedentes_medicos.GetAntecedenteMedicoServiceID(id)
	}
}

func AntecedentesMedicosPostService(data interface{}) error {
	return antecedentes_medicos.PostAntecedenteMedicoService(data)
}

func AntecedentesMedicosPutService(data interface{}) error {
	return antecedentes_medicos.PutAntecedenteMedicoService(data)
}

func AntecedentesMedicosDeleteService(id int) error {
	return antecedentes_medicos.DeleteAntecedenteMedicoService(id)
}