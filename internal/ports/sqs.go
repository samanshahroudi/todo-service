package ports

import "github.com/samanshahroudi/todo-service/internal/domain"

type SQSService interface {
	SendMessage(todo domain.TodoItem) error
}
