# Go User Management API - Dockerized

A simple REST API built with Go that persists user data to disk. Each user is stored as a JSON file on the filesystem.

## Features

- RESTful API for user management
- Persistent storage (file-based)
- Dockerized application
- No external dependencies

## Docker Hub Image

**Direct Link:** `https://hub.docker.com/r/ebeed/go-user-api`

## API Endpoints

### Add User (POST)

```bash
POST http://localhost:3000/users
Content-Type: application/json

{
  "Name": "Mohammed",
  "Age": 7
}
```

### Get User (GET)

```bash
GET http://localhost:3000/users?id=1
```

## Running with Docker

### Pull and Run from Docker Hub

```bash
docker pull ebeed/go-user-api:latest
docker run -p 3000:3000 ebeed/go-user-api:latest
```

### Build Locally

```bash
docker build -t go-user-api .
docker run -p 3000:3000 go-user-api
```

## Running Without Docker

### Prerequisites

- Go 1.21 or higher

### Steps

```bash
go mod download
go run main.go
```

## Testing the API

### Using curl

**Add a user:**

```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"Name": "Mohammed", "Age": 7}'
```

**Get a user:**

```bash
curl http://localhost:3000/users?id=1
```

### Using the provided test clients

**Add user:**

```bash
cd consumers
go run add_user.go
```

**Get user:**

```bash
cd consumers
go run get_user.go
```

## Project Structure

```
.
├── main.go              # Application entry point
├── go.mod               # Go module file
├── Dockerfile           # Docker configuration
├── README.md            # This file
├── controllers/
│   ├── fronts.go       # Route registration
│   └── user.go         # User controller
├── models/
│   ├── user.go         # User model with file persistence
│   └── address.go      # Address model
├── consumers/          # Test clients
│   ├── add_user.go
│   └── get_user.go
└── users_saved/        # Directory for user JSON files
```

## How It Works

1. When a user is created via POST request, a JSON file is created in `users_saved/{id}.txt`
2. Each file contains the complete user object in JSON format
3. When querying for a user, the application reads from the corresponding file
4. User data persists even after server restarts

## Docker Configuration

- Base image: `golang:1.21-alpine` (builder stage)
- Final image: `alpine:latest` (minimal size)
- Exposed port: `3000`
- Volume mount: Optional for persistent data across container restarts

## Author

Created as part of a Go and Docker learning assignment.

## License

MIT
