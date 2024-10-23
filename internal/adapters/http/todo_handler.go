package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samanshahroudi/todo-service/internal/usecases"
)

type TodoHandler struct {
	UseCase *usecases.TodoUseCase
}

func NewTodoHandler(useCase *usecases.TodoUseCase) *TodoHandler {
	return &TodoHandler{UseCase: useCase}
}

func (h *TodoHandler) CreateTodoItem(c *gin.Context) {
	var req struct {
		Description string    `json:"description" binding:"required"`
		DueDate     time.Time `json:"dueDate" binding:"required"`
		FileID      string    `json:"fileId"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.UseCase.CreateTodo(req.Description, req.DueDate, req.FileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
