package interfaz

import (
	"claseobrera/domains/object/files"
)

type FilesInterface func() ([]files.File, error)

func FilesGetService() FilesInterface {
	return func() ([]files.File, error) {
		return files.GetFilesService()
	}
}

func FilesGetServiceID(id string) FilesInterface {
	return func() ([]files.File, error) {
		return files.GetFileServiceID(id)
	}
}

func FilesPostService(data interface{}) error {
	return files.PostFileService(data)
}

func FilesPutService(data interface{}) error {
	return files.PutFileService(data)
}

func FilesDeleteService(id string) error {
	return files.DeleteFileService(id)
}