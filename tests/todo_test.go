package tests

import (
	"bytes"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/samanshahroudi/todo-service/internal/domain"
	"github.com/samanshahroudi/todo-service/internal/usecases"
	"github.com/samanshahroudi/todo-service/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	mockRepo := mocks.NewTodoRepository(t)
	mockS3 := mocks.NewS3Service(t)
	mockSQS := mocks.NewSQSService(t)

	todoUseCase := usecases.NewTodoUseCase(mockRepo, mockS3, mockSQS)

	description := "Test Todo"
	dueDate := time.Now().Add(24 * time.Hour)
	fileID := "test-file-id"

	expected := &domain.TodoItem{
		Description: description,
		DueDate:     dueDate,
		FileID:      fileID,
	}

	mockRepo.On("Create", mock.AnythingOfType("*domain.TodoItem")).Return(nil)
	mockSQS.On("SendMessage", mock.AnythingOfType("domain.TodoItem")).Return(nil)

	createdTodo, err := todoUseCase.CreateTodo(description, dueDate, fileID)
	assert.NoError(t, err)
	assert.Equal(t, expected.Description, createdTodo.Description)
	assert.Equal(t, expected.DueDate, createdTodo.DueDate)
	assert.Equal(t, expected.FileID, createdTodo.FileID)
}

func BenchmarkCreateTodo(b *testing.B) {
	mockRepo := mocks.NewTodoRepository(b)
	mockS3 := mocks.NewS3Service(b)
	mockSQS := mocks.NewSQSService(b)

	todoUseCase := usecases.NewTodoUseCase(mockRepo, mockS3, mockSQS)

	description := "Benchmark Todo"
	dueDate := time.Now().Add(24 * time.Hour)
	fileID := "benchmark-file-id"

	mockRepo.On("Create", mock.AnythingOfType("*domain.TodoItem")).Return(nil)
	mockSQS.On("SendMessage", mock.AnythingOfType("domain.TodoItem")).Return(nil)

	for i := 0; i < b.N; i++ {
		_, err := todoUseCase.CreateTodo(description, dueDate, fileID)
		if err != nil {
			b.Fatalf("CreateTodo failed: %v", err)
		}
	}
}

func TestUploadFile(t *testing.T) {
	mockS3 := mocks.NewS3Service(t)

	fileContent := "Test file content"
	fileReader := bytes.NewReader([]byte(fileContent))
	fileID := "unique-file-id"

	mockS3.On("UploadFile", fileID, fileReader).Return(fileID, nil)

	uploadedFileID, err := mockS3.UploadFile(fileID, fileReader)
	assert.NoError(t, err)
	assert.Equal(t, fileID, uploadedFileID)
}

func BenchmarkUploadFile(b *testing.B) {
	mockS3 := mocks.NewS3Service(b)

	fileContent := "Benchmark file content"
	fileReader := bytes.NewReader([]byte(fileContent))
	fileID := "benchmark-file-id"

	mockS3.On("UploadFile", fileID, mock.AnythingOfType("*bytes.Reader")).Return(fileID, nil)

	for i := 0; i < b.N; i++ {
		_, err := mockS3.UploadFile(fileID, fileReader)
		if err != nil {
			b.Fatalf("UploadFile failed: %v", err)
		}

		// Reset the reader to the start of the file for the next iteration
		fileReader.Seek(0, 0)
	}
}

func TestSQSSendMessage(t *testing.T) {
	mockSQS := mocks.NewSQSService(t)

	todo := domain.TodoItem{
		Description: "Test Todo",
		DueDate:     time.Now(),
		FileID:      "file-id",
	}

	mockSQS.On("SendMessage", todo).Return(nil)

	err := mockSQS.SendMessage(todo)
	assert.NoError(t, err)
}

func BenchmarkSQSSendMessage(b *testing.B) {
	mockSQS := mocks.NewSQSService(b)

	todo := domain.TodoItem{
		Description: "Benchmark Todo",
		DueDate:     time.Now(),
		FileID:      "benchmark-file-id",
	}

	mockSQS.On("SendMessage", mock.AnythingOfType("domain.TodoItem")).Return(nil)

	for i := 0; i < b.N; i++ {
		err := mockSQS.SendMessage(todo)
		if err != nil {
			b.Fatalf("SendMessage failed: %v", err)
		}
	}
}
