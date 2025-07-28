# Simple Go CRUD API

A beginner-friendly REST API built with Go that demonstrates basic CRUD (Create, Read, Update, Delete) operations for managing users.

## 🚀 Features

- **Simple HTTP Server** using Go's built-in `net/http` package
- **In-memory storage** (no database required)
- **JSON API** with proper HTTP status codes
- **Clean project structure** for learning Go basics
- **Pre-loaded sample data** for immediate testing

## 📁 Project Structure

```
simple-crud-api/
├── main.go                 # Entry point and HTTP server setup
├── models/
│   └── user.go            # User data structures
├── handlers/
│   └── user_handler.go    # HTTP request handlers
├── storage/
│   └── memory_storage.go  # In-memory data storage
├── go.mod                 # Go module file
└── README.md             # This file
```

## 🛠️ Prerequisites

- Go 1.21 or higher installed on your system
- Basic understanding of HTTP methods (GET, POST, PUT, DELETE)

## 📦 Installation

1. **Clone or create the project:**
   ```bash
   mkdir simple-crud-api
   cd simple-crud-api
   ```

2. **Create the files** as shown in the project structure above

3. **Initialize Go module:**
   ```bash
   go mod init github.com/yourusername/crud-api
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 🔗 API Endpoints

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| GET | `/health` | Health check | None |
| GET | `/users` | Get all users | None |
| GET | `/users/{id}` | Get user by ID | None |
| POST | `/users` | Create new user | `{"name":"John","email":"john@example.com"}` |
| PUT | `/users/{id}` | Update user | `{"name":"John Updated","email":"john@example.com"}` |
| DELETE | `/users/{id}` | Delete user | None |

## 📋 API Examples

### 1. Health Check
```bash
curl http://localhost:8080/health
```
**Response:**
```json
{"status":"healthy"}
```

### 2. Get All Users
```bash
curl http://localhost:8080/users
```
**Response:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2024-01-15T10:30:00Z"
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane@example.com",
    "created_at": "2024-01-15T10:30:00Z"
  }
]
```

### 3. Get User by ID
```bash
curl http://localhost:8080/users/1
```
**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-15T10:30:00Z"
}
```

### 4. Create New User
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com"}'
```
**Response:**
```json
{
  "id": 3,
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "created_at": "2024-01-15T10:35:00Z"
}
```

### 5. Update User
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john.updated@example.com"}'
```
**Response:**
```json
{
  "id": 1,
  "name": "John Updated",
  "email": "john.updated@example.com",
  "created_at": "2024-01-15T10:30:00Z"
}
```

### 6. Delete User
```bash
curl -X DELETE http://localhost:8080/users/1
```
**Response:** `204 No Content` (empty response body)

## 🔍 Code Walkthrough

### main.go
- Sets up HTTP server and routes
- Creates storage and handlers
- Starts server on port 8080

### models/user.go
- Defines `User` struct with JSON tags
- Defines `CreateUserRequest` for API input

### storage/memory_storage.go
- Implements in-memory user storage
- Provides CRUD operations
- Includes sample data initialization

### handlers/user_handler.go
- Handles HTTP requests and routing
- Converts between JSON and Go structs
- Returns appropriate HTTP status codes

## 🧪 Testing

You can test the API using:

1. **curl** (as shown in examples above)
2. **Postman** or **Insomnia** (import the endpoints)
3. **Go test files** (you can create test files later)

## 🎯 Learning Objectives

After working with this code, you'll understand:

- ✅ Basic Go project structure
- ✅ HTTP server setup with `net/http`
- ✅ JSON handling with struct tags
- ✅ Method routing based on HTTP verbs
- ✅ Error handling and HTTP status codes
- ✅ In-memory data storage with maps
- ✅ Go interfaces and method receivers

## 🚀 Next Steps

Once you're comfortable with this basic version, consider:

1. **Add input validation** (email format, required fields)
2. **Add a real database** (PostgreSQL, MySQL, SQLite)
3. **Add middleware** (logging, CORS, authentication)
4. **Use a router library** (Gorilla Mux, Gin, Echo)
5. **Add tests** (unit tests, integration tests)
6. **Add configuration** (environment variables, config files)
7. **Add documentation** (Swagger/OpenAPI)

## 🤝 Contributing

This is a learning project! Feel free to:
- Add features
- Improve error handling
- Add tests
- Refactor code structure

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

---

**Happy Learning! 🎉**

Built with ❤️ and Go