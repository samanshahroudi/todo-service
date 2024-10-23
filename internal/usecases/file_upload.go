package usecases

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"github.com/samanshahroudi/todo-service/internal/ports"
)

type FileUploadUseCase struct {
	S3Service    ports.S3Service
	MaxFileSize  int64
	AllowedTypes []string
}

func NewFileUploadUseCase(s3Service ports.S3Service) *FileUploadUseCase {
	return &FileUploadUseCase{
		S3Service:    s3Service,
		MaxFileSize:  5 << 20, // Max file size: 5MB
		AllowedTypes: []string{"image/jpeg", "image/png", "text/plain"},
	}
}

func (u *FileUploadUseCase) ValidateFile(file *multipart.FileHeader) error {
	// Check file size
	if file.Size > u.MaxFileSize {
		return errors.New("file size exceeds 5MB")
	}

	// Validate content type
	openedFile, err := file.Open()
	if err != nil {
		return errors.New("cannot open file")
	}
	defer func(openedFile multipart.File) {
		err := openedFile.Close()
		if err != nil {
			panic(err)
		}
	}(openedFile)

	fileHeader := make([]byte, 512)
	if _, err := openedFile.Read(fileHeader); err != nil {
		return errors.New("cannot read file")
	}

	contentType := http.DetectContentType(fileHeader)
	isAllowed := false
	for _, t := range u.AllowedTypes {
		if t == contentType {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return errors.New("unsupported file type")
	}

	return nil
}

func (u *FileUploadUseCase) UploadFile(fileID string, file multipart.File) (string, error) {
	// Upload the file using the S3 service
	return u.S3Service.UploadFile(fileID, file)
}

func (u *FileUploadUseCase) GenerateFileID() string {
	// Generate a unique file ID
	return uuid.New().String()
}
