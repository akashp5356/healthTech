package service

import (
	"errors"
	"fmt"
	"healtech-backend/server/internal/config"
	"healtech-backend/server/internal/models"
	"healtech-backend/server/internal/repository"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type DocumentService struct {
	Config *config.Config
}

func NewDocumentService(cfg *config.Config) *DocumentService {
	return &DocumentService{Config: cfg}
}
func (s *DocumentService) ListDocumentsService(userID int) ([]models.DocumentDetails, error) {
	// cacheKey := fmt.Sprintf("docs:user:%d", userID)
	// ctx := context.Background()

	// // Try Cache
	// val, err := repository.RDB.Get(ctx, cacheKey).Result()
	// if err == nil {
	// 	var docs []models.DocumentDetails
	// 	if err := json.Unmarshal([]byte(val), &docs); err == nil {
	// 		return docs, nil
	// 	}
	// }
	// Fetch from DB
	docs, err := repository.ListDocumentsRepo(userID)
	if err != nil {
		return nil, err
	}

	// // Set Cache
	// data, _ := json.Marshal(docs)
	// repository.RDB.Set(ctx, cacheKey, data, 10*time.Minute)

	return docs, nil
}

func (s *DocumentService) UploadDocumentService(file *multipart.FileHeader, userID int, docTypeID int, description string) (*models.DocumentDetails, error) {
	// 1. Validation
	if file.Size > 10*1024*1024 { // 10MB
		return nil, errors.New("file too large")
	}
	if filepath.Ext(file.Filename) != ".pdf" {
		// Strict check: Read first 512 bytes for MIME type if needed, but keeping it simple for PoC
		return nil, errors.New("only PDF files are allowed")
	}

	// 2. Prepare Storage
	if _, err := os.Stat(s.Config.UploadDir); os.IsNotExist(err) {
		os.MkdirAll(s.Config.UploadDir, os.ModePerm)
	}

	newFilename := uuid.New().String() + ".pdf"
	storagePath := filepath.Join(s.Config.UploadDir, newFilename)

	// 3. Save to Disk
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(storagePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	// 4. Save to DB
	doc := &models.DocumentDetails{
		UserID:           userID,
		DocumentID:       docTypeID,
		FilePath:         storagePath,
		OriginalFilename: file.Filename,
		FileSize:         file.Size,
		Description:      description,
	}

	id, err := repository.CreateDocumentRepo(doc)
	if err != nil {
		os.Remove(storagePath) // Cleanup
		return nil, err
	}
	doc.ID = int(id)
	doc.UploadedAt = time.Now()

	return doc, nil
}

func (s *DocumentService) GetDocumentPathService(docID, userID int) (string, string, error) {
	doc, err := repository.GetDocumentByIDRepo(docID)
	if err != nil {
		return "", "", err
	}

	if doc.UserID != userID {
		return "", "", errors.New("unauthorized")
	}

	return doc.FilePath, doc.OriginalFilename, nil
}

func (s *DocumentService) DeleteDocumentService(docID, userID int) error {
	doc, err := repository.GetDocumentByIDRepo(docID)
	if err != nil {
		return err
	}

	if doc.UserID != userID {
		return errors.New("unauthorized")
	}

	// Delete from DB
	if err := repository.DeleteDocumentRepo(docID); err != nil {
		return err
	}
	fmt.Println("-->", doc.FilePath)
	os.Remove(doc.FilePath)

	// s.deleteFrpmCache(userID)
	return nil
}
