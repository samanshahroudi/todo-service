package http

import (
	"github.com/gin-gonic/gin"
	"github.com/samanshahroudi/todo-service/internal/usecases"
	"mime/multipart"
	"net/http"
)

type FileHandler struct {
	UseCase *usecases.FileUploadUseCase
}

func NewFileHandler(useCase *usecases.FileUploadUseCase) *FileHandler {
	return &FileHandler{
		UseCase: useCase,
	}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	if err := h.UseCase.ValidateFile(file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open file"})
		return
	}
	defer func(openedFile multipart.File) {
		err := openedFile.Close()
		if err != nil {

		}
	}(openedFile)

	// Generate a unique file ID using the use case
	fileID := h.UseCase.GenerateFileID()

	// Upload the file using the use case
	_, err = h.UseCase.UploadFile(fileID, openedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fileId": fileID})
}
