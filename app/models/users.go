package models

import (
	"claseobrera/app/repository"
	"claseobrera/sessions/interfaz"
)

func GetUsers() (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		org := interfaz.UsersService()
		result, err := org()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func GetUsersID(id int) (interface{}, error) {
	return repository.FetchDomainService(func() (interface{}, error) {
		orgID := interfaz.UsersIDService(id)
		result, err := orgID()
		if err != nil {
			return nil, err
		}
		return result, nil
	})
}

func PostUsers(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.UsersPostService(data)
	})
}

func PutUsers(data interface{}) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.UsersPutService(data)
	})
}

func DeleteUsersID(id int) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.UsersDeleteService(id)
	})
}
