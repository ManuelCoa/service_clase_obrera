package models

import (
	"claseobrera/app/repository"
	"claseobrera/domains/interfaz"
)

func GetFiles() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.FilesGetService()
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetFileID(id string) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		data := interfaz.FilesGetServiceID(id)
		result, err := data()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostFile(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.FilesPostService(data)
	})
}

func PutFile(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.FilesPutService(data)
	})
}

func DeleteFile(id string) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.FilesDeleteService(id)
	})
}