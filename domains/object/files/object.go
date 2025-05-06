package files

import (
	"time"
)

type File struct {
	ID          string   `json:"id_files"`
	OriginalName string    `json:"original_name"`
	MimeType    string    `json:"mime_type"`
	UploadDate  time.Time `json:"upload_date"`
	StoragePath string    `json:"storage_path"`
}