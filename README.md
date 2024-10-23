### 1. Clone the repository

```bash
git clone <repository-url>
cd todo-service
```

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

### License
This project is licensed under the MIT License.
