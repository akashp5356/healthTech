package handler

import (
	"fmt"
	"healtech-backend/server/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DocService  *service.DocumentService
	AuthService *service.AuthService
}

func NewHandler(docService *service.DocumentService, authService *service.AuthService) *Handler {
	return &Handler{
		DocService:  docService,
		AuthService: authService,
	}
}

func (h *Handler) ListDocuments(c *gin.Context) {
	userID := c.GetInt("userID")

	docs, err := h.DocService.ListDocumentsService(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"documents": docs})
}

func (h *Handler) UploadDocument(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	description := c.PostForm("description")
	docTypeIDStr := c.DefaultPostForm("document_type_id", "1") // Default to 1 (prescription)
	docTypeID, _ := strconv.Atoi(docTypeIDStr)

	userID := c.GetInt("userID")

	doc, err := h.DocService.UploadDocumentService(file, userID, docTypeID, description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, doc)
}

func (h *Handler) DownloadDocument(c *gin.Context) {
	userID := c.GetInt("userID")
	docID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	path, filename, err := h.DocService.GetDocumentPathService(docID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found or access denied"})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.File(path)
}

func (h *Handler) DeleteDocument(c *gin.Context) {
	userID := c.GetInt("userID")
	docID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	if err := h.DocService.DeleteDocumentService(docID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted"})
}
