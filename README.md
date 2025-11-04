# go-healthcare

# Go Health API

A simple REST API built with Go and the Gin framework. This project demonstrates basic CRUD operations, RESTful API design, and Go best practices.

## 🎯 Purpose

This project was built as a learning exercise to:
- Understand Go syntax and idioms
- Learn the Gin web framework
- Practice REST API design in Go
- Demonstrate Go skills for the Corti interview

## 🚀 Features

- ✅ Health check endpoint
- ✅ Complete CRUD operations for users
- ✅ RESTful API design
- ✅ JSON request/response handling
- ✅ Error handling
- ✅ In-memory data storage
- ✅ Clean code structure

## 📋 Prerequisites

- Go 1.21 or higher (tested with Go 1.25.3)
- Git (optional, for cloning)

## 🔧 Installation

### 1. Install Go

Download and install Go from [https://go.dev/dl/](https://go.dev/dl/)

Verify installation:
```bash
go version
```

### 2. Clone or Download This Repository

```bash
git clone https://github.com/Kether180/go-health-api.git
cd go-health-api
```

Or simply download the files and navigate to the directory.

### 3. Install Dependencies

If you encounter checksum mismatch errors, regenerate the go.sum file:

```bash
# Remove old go.sum and regenerate
rm go.sum
go mod tidy
```

Then download dependencies:

```bash
go mod download
```

This will download the Gin framework and its dependencies.

**Note:** If you see `SECURITY ERROR` related to checksum mismatches, it's likely due to Go version differences. The solution above will regenerate checksums for your specific Go version.
o
## 🏃 Running the API

### Option 1: Run Directly

Start the server:

```bash
go run main.go
```

### Option 2: Build and Run

Build an executable:

```bash
go build -o health-api.exe main.go
```

Then run it:

```bash
./health-api.exe
```

You should see:
```
[GIN-debug] Listening and serving HTTP on :8080
```

The API is now running at `http://localhost:8080`

## 📡 API Endpoints

### Health Check

Check if the service is running:

```bash
GET /health
```

**Example:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "healthy",
  "message": "Service is running",
  "version": "1.0.0"
}
```

---

### Get All Users

Retrieve all users:

```bash
GET /api/v1/users
```

**Example:**
```bash
curl http://localhost:8080/api/v1/users
```

**Response:**
```json
{
  "count": 3,
  "data": [
    {
      "id": 1,
      "name": "Alice Johnson",
      "email": "alice@example.com"
    },
    {
      "id": 2,
      "name": "Bob Smith",
      "email": "bob@example.com"
    },
    {
      "id": 3,
      "name": "Carol Davis",
      "email": "carol@example.com"
    }
  ]
}
```

---

### Get User by ID

Retrieve a specific user:

```bash
GET /api/v1/users/:id
```

**Example:**
```bash
curl http://localhost:8080/api/v1/users/1
```

**Response:**
```json
{
  "data": {
    "id": 1,
    "name": "Alice Johnson",
    "email": "alice@example.com"
  }
}
```

---

### Create New User

Create a new user:

```bash
POST /api/v1/users
Content-Type: application/json

{
  "name": "David Wilson",
  "email": "david@example.com"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"David Wilson","email":"david@example.com"}'
```

**Response:**
```json
{
  "message": "User created successfully",
  "data": {
    "id": 4,
    "name": "David Wilson",
    "email": "david@example.com"
  }
}
```

---

### Update User

Update an existing user:

```bash
PUT /api/v1/users/:id
Content-Type: application/json

{
  "name": "David Wilson Updated",
  "email": "david.updated@example.com"
}
```

**Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/4 \
  -H "Content-Type: application/json" \
  -d '{"name":"David Wilson Updated","email":"david.updated@example.com"}'
```

**Response:**
```json
{
  "message": "User updated successfully",
  "data": {
    "id": 4,
    "name": "David Wilson Updated",
    "email": "david.updated@example.com"
  }
}
```

---

### Delete User

Delete a user:

```bash
DELETE /api/v1/users/:id
```

**Example:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/4
```

**Response:**
```json
{
  "message": "User deleted successfully"
}
```

## 🧪 Testing with Postman

1. Download and install [Postman](https://www.postman.com/downloads/)
2. Import the provided Postman collection (if available)
3. Or manually create requests for each endpoint above

## 📁 Project Structure

```
go-health-api/
├── main.go           # Main application file with all API logic
├── go.mod            # Go module dependencies
├── go.sum            # Dependency checksums (auto-generated)
├── README.md         # Project documentation
├── SETUP_GUIDE.md    # Detailed setup instructions
├── TESTING.md        # Testing guide
├── .gitignore        # Git ignore file
└── health-api.exe    # Compiled binary (after building)
```

## 🎓 What I Learned

Building this project helped me understand:

1. **Go Syntax & Idioms**
   - Package structure
   - Struct definitions with JSON tags
   - Error handling with `if err != nil`
   - Slice operations

2. **Gin Framework**
   - Router setup with `gin.Default()`
   - Route grouping with `r.Group()`
   - HTTP method handlers (GET, POST, PUT, DELETE)
   - JSON binding with `c.ShouldBindJSON()`
   - Response formatting with `c.JSON()`

3. **REST API Design**
   - RESTful endpoint structure
   - HTTP status codes (200, 201, 400, 404)
   - Request/response patterns
   - Error response formatting

4. **Go Best Practices**
   - Explicit error handling
   - Clear function naming
   - JSON struct tags
   - HTTP status code usage

## 🔄 Next Steps for Improvement

If I had more time, I would add:

- [ ] **Thread Safety**: Add mutex locks for concurrent access to in-memory data
- [ ] **Input Validation**: Email format validation and improved input sanitization
- [ ] **Database Integration**: PostgreSQL or SQLite for persistent storage
- [ ] **Unit Tests**: Comprehensive tests using Go's `testing` package
- [ ] **Middleware**: Logging, authentication, and request/response middleware
- [ ] **API Documentation**: Swagger/OpenAPI specification
- [ ] **Docker**: Containerization for easy deployment
- [ ] **Environment Configuration**: .env file support for configurable settings
- [ ] **Rate Limiting**: Protection against abuse
- [ ] **CORS Support**: Cross-origin resource sharing for web clients

## 🎯 Interview Talking Points

**Why Go?**
- Compiled language with excellent performance
- Built-in concurrency with goroutines
- Simple, clean syntax
- Great for backend APIs and microservices
- Single binary deployment (no dependencies)

**What I Learned:**
- Go's explicit error handling forces you to think about edge cases
- The Gin framework makes REST API development fast and clean
- Go's struct tags make JSON serialization simple
- The standard library is powerful and well-designed

**How This Relates to Corti:**
- Real-time health checks are essential for production services
- CRUD operations are fundamental to any backend system
- Error handling is critical for healthcare APIs where reliability matters

## 📚 Resources Used

- [Official Go Documentation](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Gin Framework Documentation](https://gin-gonic.com/docs/)
- [Effective Go](https://go.dev/doc/effective_go)


