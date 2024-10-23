package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	AWSAddress string
	S3Bucket   string
	S3Region   string

	SQSQueueURL string
	SQSRegion   string
}

func LoadConfig() Config {
	return Config{
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "password"),
		DBName:      getEnv("DB_NAME", "todo_db"),
		AWSAddress:  getEnv("AWS_ADDRESS", "http://localhost:4566"),
		S3Bucket:    getEnv("S3_BUCKET", "todo-bucket.s3"),
		S3Region:    getEnv("S3_REGION", "us-east-1"),
		SQSQueueURL: getEnv("SQS_QUEUE_URL", "http://localhost:4566/000000000000/todo-queue"),
		SQSRegion:   getEnv("SQS_REGION", "us-east-1"),
	}
}

// Helper function to get an environment variable or use a default value if not set
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
