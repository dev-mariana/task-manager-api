# Task Manager API

A RESTful API built with Go and Gin framework for managing tasks. This API provides endpoints to create, read, update, and delete tasks with status tracking.

## Features

- ✅ Create, read, update, and delete tasks
- ✅ Task status management (Pending, In Progress, Completed)
- ✅ PostgreSQL database integration with GORM
- ✅ Swagger/OpenAPI documentation
- ✅ Docker Compose setup for database
- ✅ Clean architecture with handlers, services, and repositories

## Tech Stack

- **Go** 1.25.3
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **PostgreSQL** - Database
- **Swagger/OpenAPI** - API documentation

## Prerequisites

- Go 1.25.3 or higher
- Docker and Docker Compose (for database)
- PostgreSQL (if not using Docker)

## Getting Started

### 1. Clone the Repository

```bash
git clone <git@github.com:dev-mariana/task-manager-api.git>
cd task-manager-api
```

### 2. Set Up the Database

Using Docker Compose (recommended):

```bash
docker-compose up -d
```

This will start a PostgreSQL container with the following default settings:

- **Host**: localhost
- **Port**: 5432
- **Username**: docker
- **Password**: docker@123
- **Database**: task-manager

### 3. Configure Environment Variables

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=

# Server Configuration
SERVER_PORT=8080
```

### 4. Install Dependencies

```bash
go mod download
```

### 5. Run the Application

```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080` (or the port specified in `SERVER_PORT`).

## API Documentation

Once the server is running, you can access the interactive Swagger documentation at:

**http://localhost:8080/swagger/index.html**

The Swagger UI provides:

- Complete API endpoint documentation
- Request/response schemas
- Interactive testing interface
- Example models and data structures

## API Endpoints

### Tasks

| Method   | Endpoint     | Description       |
| -------- | ------------ | ----------------- |
| `POST`   | `/tasks`     | Create a new task |
| `GET`    | `/tasks`     | Get all tasks     |
| `GET`    | `/tasks/:id` | Get a task by ID  |
| `PATCH`  | `/tasks/:id` | Update a task     |
| `DELETE` | `/tasks/:id` | Delete a task     |

### Request/Response Examples

#### Create Task

**Request:**

```bash
POST /tasks
Content-Type: application/json

{
  "id": "unique-task-id",
  "title": "Complete project documentation",
  "description": "Write comprehensive README and API docs",
  "status": "Pending"
}
```

**Response:**

```json
{
  "id": "unique-task-id",
  "title": "Complete project documentation",
  "description": "Write comprehensive README and API docs",
  "status": "Pending",
  "created_at": "2025-11-09T15:00:00Z",
  "updated_at": "2025-11-09T15:00:00Z"
}
```

#### Get All Tasks

**Request:**

```bash
GET /tasks
```

**Response:**

```json
[
  {
    "id": "task-1",
    "title": "Task 1",
    "description": "Description 1",
    "status": "Pending",
    "created_at": "2025-11-09T15:00:00Z",
    "updated_at": "2025-11-09T15:00:00Z"
  },
  {
    "id": "task-2",
    "title": "Task 2",
    "description": "Description 2",
    "status": "In Progress",
    "created_at": "2025-11-09T15:00:00Z",
    "updated_at": "2025-11-09T15:00:00Z"
  }
]
```

#### Update Task

**Request:**

```bash
PATCH /tasks/:id
Content-Type: application/json

{
  "title": "Updated title",
  "description": "Updated description",
  "status": "Completed"
}
```

**Response:**

```json
{
  "id": "unique-task-id",
  "title": "Updated title",
  "description": "Updated description",
  "status": "Completed",
  "created_at": "2025-11-09T15:00:00Z",
  "updated_at": "2025-11-09T15:10:00Z"
}
```

**Note:** All fields in the update request are optional.

#### Delete Task

**Request:**

```bash
DELETE /tasks/:id
```

**Response:**

```json
{
  "message": "Task deleted successfully"
}
```

## Task Status

Tasks can have one of the following statuses:

- `Pending` - Task is not started
- `In Progress` - Task is currently being worked on
- `Completed` - Task is finished

## Project Structure

```
task-manager-api/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── database.go         # Database configuration
│   ├── entities/
│   │   └── task.go             # Task entity and DTOs
│   ├── handlers/
│   │   ├── create-task.handler.go
│   │   ├── delete-task.handler.go
│   │   ├── get-all-tasks.handler.go
│   │   ├── get-task.handler.go
│   │   └── update-task.handler.go
│   ├── repositories/
│   │   ├── task.interface.go
│   │   └── task.repository.go
│   └── services/
│       ├── create-task.service.go
│       ├── delete-task.service.go
│       ├── get-all-tasks.service.go
│       ├── get-task.service.go
│       └── update-task.service.go
├── routes/
│   └── routes.go               # Route definitions
├── docs/
│   ├── docs.go                 # Generated Swagger docs
│   ├── swagger.json
│   └── swagger.yaml
├── docker-compose.yaml         # Docker Compose configuration
├── go.mod
├── go.sum
└── README.md
```

## Regenerating Swagger Documentation

If you modify the API endpoints or add new ones, regenerate the Swagger documentation:

```bash
swag init -g cmd/main.go
```

Make sure you have `swag` installed:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## License

MIT
