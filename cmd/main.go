package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samanshahroudi/todo-service/internal/adapters/http"
	"github.com/samanshahroudi/todo-service/internal/adapters/repository"
	"github.com/samanshahroudi/todo-service/internal/adapters/s3"
	"github.com/samanshahroudi/todo-service/internal/adapters/sqs"
	"github.com/samanshahroudi/todo-service/internal/domain"
	"github.com/samanshahroudi/todo-service/internal/usecases"
	"github.com/samanshahroudi/todo-service/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	awsSQS "github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	loadConfig := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		loadConfig.DBHost, loadConfig.DBPort, loadConfig.DBUser, loadConfig.DBPassword, loadConfig.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&domain.TodoItem{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(loadConfig.S3Region),
	)
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	s3Client := awsS3.NewFromConfig(awsCfg, func(o *awsS3.Options) {
		o.BaseEndpoint = &loadConfig.AWSAddress
	})

	s3Service := s3.NewS3Adapter(s3Client, loadConfig.S3Bucket)

	sqsClient := awsSQS.NewFromConfig(awsCfg, func(o *awsSQS.Options) {
		o.BaseEndpoint = &loadConfig.AWSAddress
	})
	sqsService := sqs.NewSQSAdapter(sqsClient, loadConfig.SQSQueueURL)

	repo := repository.NewPostgresRepository(db)

	todoUseCase := usecases.NewTodoUseCase(repo, s3Service, sqsService)

	// Initialize HTTP Handlers
	todoHandler := http.NewTodoHandler(todoUseCase)

	fileHandlerUseCase := usecases.NewFileUploadUseCase(s3Service)
	fileHandler := http.NewFileHandler(fileHandlerUseCase)

	router := gin.Default()

	// Define routes
	router.POST("/upload", fileHandler.UploadFile)
	router.POST("/todo", todoHandler.CreateTodoItem)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
