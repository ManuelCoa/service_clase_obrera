package files

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectFiles() (*sql.Rows, error) {
	query := "SELECT id_files, original_name, mime_type, upload_date, storage_path FROM files"
	return repository.FetchRows(query)
}

func selectFileID(id string) (*sql.Rows, error) {
	query := "SELECT id_files, original_name, mime_type, upload_date, storage_path FROM files WHERE id_files = ?"
	return repository.FetchRows(query, id)
}

func insertFile(data File) (sql.Result, error) {
	query := "INSERT INTO files (id_files, original_name, mime_type, upload_date, storage_path) VALUES (?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, 
		data.ID,
		data.OriginalName,
		data.MimeType,
		data.UploadDate,
		data.StoragePath,
	)
}

func updateFile(data File) (sql.Result, error) {
	query := "UPDATE files SET original_name = ?, mime_type = ?, storage_path = ? WHERE id_files = ?"
	return repository.ExecuteQuery(query, 
		data.OriginalName,
		data.MimeType,
		data.StoragePath,
		data.ID,
	)
}

func deleteFile(id string) (sql.Result, error) {
	query := "DELETE FROM files WHERE id_files = ?"
	return repository.ExecuteQuery(query, id)
}