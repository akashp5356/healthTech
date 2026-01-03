package models

import "time"

type RoleMaster struct {
	ID        int       `json:"id" db:"id"`
	RoleName  string    `json:"role_name" db:"role_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type LoginMaster struct {
	ID           int       `json:"id" db:"id"`
	RegisterID   int       `json:"register_id" db:"register_id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password_hash"`
	RoleID       int       `json:"role_id" db:"role_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type DocumentMaster struct {
	ID           int       `json:"id" db:"id"`
	DocumentType string    `json:"document_type" db:"document_type"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type RegisterDetails struct {
	ID        int       `json:"id" db:"id"`
	FullName  string    `json:"full_name" db:"full_name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type DocumentDetails struct {
	ID               int       `json:"id" db:"id"`
	UserID           int       `json:"user_id" db:"user_id"`
	DocumentID       int       `json:"document_id" db:"document_id"`
	FilePath         string    `json:"-" db:"file_path"`
	OriginalFilename string    `json:"filename" db:"original_filename"`
	FileSize         int64     `json:"file_size" db:"file_size"`
	Description      string    `json:"description" db:"description"`
	UploadedAt       time.Time `json:"uploaded_at" db:"uploaded_at"`
	DocumentType     string    `json:"document_type,omitempty" `
}
