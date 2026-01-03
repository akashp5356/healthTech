package service

import (
	"healtech-backend/server/internal/config"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mocking config
var mockConfig = &config.Config{
	UploadDir: "./uploads_test",
	JWTSecret: "testsecret",
}

func TestUploadDocument_FileSizeExceeded(t *testing.T) {
	svc := NewDocumentService(mockConfig)

	// Create a dummy file header with large size
	fileHeader := &multipart.FileHeader{
		Filename: "test.pdf",
		Size:     11 * 1024 * 1024, // 11 MB
	}

	_, err := svc.UploadDocumentService(fileHeader, 1, 1, "desc")
	assert.Error(t, err)
	assert.Equal(t, "file too large", err.Error())
}

func TestUploadDocument_InvalidExtension(t *testing.T) {
	svc := NewDocumentService(mockConfig)

	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg", // Invalid
		Size:     1024,
	}

	_, err := svc.UploadDocumentService(fileHeader, 1, 1, "desc")
	assert.Error(t, err)
	assert.Equal(t, "only PDF files are allowed", err.Error())
}

func TestAuthService_ValidateToken_Invalid(t *testing.T) {
	svc := NewAuthService(mockConfig)

	_, err := svc.ValidateToken("invalid.token.string")
	assert.Error(t, err)
}

func TestConfig_Defaults(t *testing.T) {
	// 3. Test Config Loading defaults
	cfg := config.LoadConfig()
	assert.Equal(t, "8080", cfg.AppPort) // Default
}
