package usecases

import (
	"time"

	"github.com/samanshahroudi/todo-service/internal/domain"
	"github.com/samanshahroudi/todo-service/internal/ports"
)

type TodoUseCase struct {
	Repo ports.TodoRepository
	S3   ports.S3Service
	SQS  ports.SQSService
}

func NewTodoUseCase(repo ports.TodoRepository, s3 ports.S3Service, sqs ports.SQSService) *TodoUseCase {
	return &TodoUseCase{
		Repo: repo,
		S3:   s3,
		SQS:  sqs,
	}
}

func (u *TodoUseCase) CreateTodo(description string, dueDate time.Time, fileID string) (*domain.TodoItem, error) {
	todo := domain.NewTodoItem(description, dueDate, fileID)
	if err := u.Repo.Create(todo); err != nil {
		return nil, err
	}

	if err := u.SQS.SendMessage(*todo); err != nil {
		return nil, err
	}

	return todo, nil
}
