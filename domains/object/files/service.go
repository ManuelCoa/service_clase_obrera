package files

func GetFilesService() ([]File, error) {
	return searchParsedFiles()
}

func GetFileServiceID(id string) ([]File, error) {
	return searchParsedFileID(id)
}

func PostFileService(data interface{}) error {
	return insertionFile(data)
}

func PutFileService(data interface{}) error {
	return putFile(data)
}

func DeleteFileService(id string) error {
	return deletionFile(id)
}