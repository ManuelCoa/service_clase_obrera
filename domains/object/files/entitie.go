package files

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *File) error {
	return rows.Scan(
		&dato.ID,
		&dato.OriginalName,
		&dato.MimeType,
		&dato.UploadDate,
		&dato.StoragePath,
	)
}

func searchParsedFiles() ([]File, error) {
	rows, err := selectFiles()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedFileID(id string) ([]File, error) {
	rows, err := selectFileID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionFile(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[File](data)
	if err != nil {
		return err
	}

	_, err = insertFile(dato)
	return err
}

func putFile(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[File](data)
	if err != nil {
		return err
	}

	_, err = updateFile(dato)
	return err
}

func deletionFile(id string) error {
	_, err := deleteFile(id)
	return err
}