package interfaz

import (
	"claseobrera/domains/object/reposo_medico"
)

type ReposoMedicoInterface func() ([]reposo_medico.ReposoMedico, error)

func ReposoMedicoGetService() ReposoMedicoInterface {
	return func() ([]reposo_medico.ReposoMedico, error) {
		return reposo_medico.GetRepososService()
	}
}

func ReposoMedicoGetServiceID(id int) ReposoMedicoInterface {
	return func() ([]reposo_medico.ReposoMedico, error) {
		return reposo_medico.GetRepososServiceID(id)
	}
}

func ReposoMedicoPostService(data interface{}) error {
	return reposo_medico.PostReposoService(data)
}

func ReposoMedicoPutService(data interface{}) error {
	return reposo_medico.PutReposoService(data)
}

func ReposoMedicoDeleteService(id int) error {
	return reposo_medico.DeleteReposoService(id)
}