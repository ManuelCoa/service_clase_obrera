package reposo_medico

func GetRepososService() ([]ReposoMedico, error) {
	return searchParsedReposos()
}

func GetRepososServiceID(id int) ([]ReposoMedico, error) {
	return searchParsedRepososID(id)
}

func PostReposoService(data interface{}) error {
	return insertionReposo(data)
}

func PutReposoService(data interface{}) error {
	return putReposo(data)
}

func DeleteReposoService(id int) error {
	return deletionReposo(id)
}