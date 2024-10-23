package ports

import "github.com/gin-gonic/gin"

type TodoHTTPPort interface {
	UploadFile(c *gin.Context)
	CreateTodoItem(c *gin.Context)
}
