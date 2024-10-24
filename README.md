# Todo Service

This project is a simple Todo management service that integrates with AWS S3 for file uploads and SQS for messaging, and uses PostgreSQL as its database. It also includes unit tests and benchmarks for Todo item creation, file uploads, and SQS messaging.

## Prerequisites

Before you can run this project, you need to have the following installed:

- Docker
- Docker Compose
- Go (if running locally)

## Getting Started

To start the project, apply migrations, and run tests, follow the steps below.

### 1. Clone the repository

```bash
git clone https://github.com/samanshahroudi/todo-service.git
cd todo-service
```

### 2. Set Up Environment Variables
Create a .env file at the root of the project and configure the environment variables:

```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=todo_db
S3_BUCKET=todo-bucket.s3
S3_REGION=us-east-1
SQS_QUEUE_URL=http://localstack:4566/000000000000/todo-queue
SQS_REGION=us-east-1
AWS_ACCESS_KEY_ID=test
AWS_SECRET_ACCESS_KEY=test
AWS_REGION=us-east-1
AWS_ADDRESS=http://localstack:4566
```

### 3. Run the Project with Docker Compose
Run the following command to start the project and its dependencies using Docker Compose:

```bash
make run
```
This will:

Start the PostgreSQL database.
Start LocalStack with S3 and SQS services.
Run the todo-service app.

### 4. Apply Migrations
If you're using GORM for auto migrations, the application will automatically apply migrations when it starts.


### 5. Run Unit Tests
To run the unit tests, use the following command:
```bash
make test
```
This command will execute all unit tests within the project.

### 6. Run Benchmarks
To run the benchmarks for Todo item creation, file uploads, and SQS messaging, use the following command:
```bash
make benchmark
```
This will run all benchmarks and output performance metrics.

### Project Structure
```
├── cmd                     # Main application entry point
├── internal                # Internal application packages
│   ├── adapters            # Adapters for S3, SQS, HTTP, etc.
│   ├── domain              # Domain entities (TodoItem, etc.)
│   ├── ports               # Interfaces for dependencies
│   └── usecases            # Application business logic
├── tests                   # Unit and benchmark tests
├── Dockerfile              # Dockerfile for building the project
├── docker-compose.yml      # Docker Compose setup for local development
├── Makefile                # Make file 
└── README.md               # Project documentation (this file)
```


### Technologies Used
* Go: Backend language
* PostgreSQL: Database
* GORM: ORM for database interactions
* AWS S3 & SQS: For file storage and messaging (via LocalStack)
* Docker: Containerization for local development

### Troubleshooting
* If you encounter issues with connecting to LocalStack or PostgreSQL, ensure that both services are up and running in Docker. You can check the logs with:
```bash
docker compose logs
```


### License
This project is licensed under the MIT License.