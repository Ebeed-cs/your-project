# Go User Management API - Dockerized

A simple REST API built with Go that persists user data to disk. Each user is stored as a JSON file on the filesystem.

## Features

- RESTful API for user management
- Persistent storage (file-based)
- Dockerized application
- No external dependencies

## Docker Hub Image

**Direct Link:** `https://hub.docker.com/r/YOUR_DOCKERHUB_USERNAME/go-user-api`

*(Replace YOUR_DOCKERHUB_USERNAME with your actual Docker Hub username)*

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
docker pull YOUR_DOCKERHUB_USERNAME/go-user-api:latest
docker run -p 3000:3000 YOUR_DOCKERHUB_USERNAME/go-user-api:latest
```

### Build Locally
```bash
docker build -t go-user-api .
docker run -p 3000:3000 go-user-api
```

## Project Structure
