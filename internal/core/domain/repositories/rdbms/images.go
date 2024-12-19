package rdbms

import "time"

type CreateImage struct {
	FileName   string    `json:"file_name"`
	S3Path     string    `db:"s3_path" json:"s3_path"`
	MimeType   string    `db:"mime_type" json:"mime_type"`
	IsActive   bool      `db:"is_active" json:"is_active"`
	UploadedBy string    `db:"uploaded_by" json:"uploaded_by"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"`
}

type Image struct {
	ID string `db:"id" json:"id"`
	CreateImage
}
type UploadImageRes struct {
	Id        string `json:"id"`
	SignedURL string `json:"signed_url"`
}
