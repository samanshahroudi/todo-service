package repository

import (
	"github.com/samanshahroudi/todo-service/internal/domain"
	"github.com/samanshahroudi/todo-service/internal/ports"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) ports.TodoRepository {
	return &PostgresRepository{
		DB: db,
	}
}

func (r *PostgresRepository) Create(todo *domain.TodoItem) error {
	return r.DB.Create(todo).Error
}
