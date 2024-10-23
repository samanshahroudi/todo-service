package ports

import "github.com/samanshahroudi/todo-service/internal/domain"

type TodoRepository interface {
	Create(todo *domain.TodoItem) error
}
