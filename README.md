# Go Health API

A simple REST API built with Go and the Gin framework. This project demonstrates basic CRUD operations, RESTful API design, and Go best practices.

## Features

- **RESTful API Design**: Clean and intuitive endpoints
- **CRUD Operations**: Create, Read, Update, and Delete health records
- **Input Validation**: Automatic validation using Gin's binding tags
- **Error Handling**: Proper HTTP status codes and error messages
- **In-Memory Storage**: Simple data storage for demonstration

## Prerequisites

- Go 1.24 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Kether180/go-healthcare.git
cd go-healthcare
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
- **GET** `/api/v1/health`
  - Returns the health status of the API
  - Response: `200 OK`
  ```json
  {
    "status": "healthy",
    "message": "Go Health API is running"
  }
  ```

### Get All Records
- **GET** `/api/v1/records`
  - Returns all health records
  - Response: `200 OK`
  ```json
  [
    {
      "id": 1,
      "patient_name": "John Doe",
      "age": 35,
      "diagnosis": "Flu",
      "treatment": "Rest and fluids"
    }
  ]
  ```

### Get Record by ID
- **GET** `/api/v1/records/:id`
  - Returns a specific health record
  - Response: `200 OK` or `404 Not Found`
  ```json
  {
    "id": 1,
    "patient_name": "John Doe",
    "age": 35,
    "diagnosis": "Flu",
    "treatment": "Rest and fluids"
  }
  ```

### Create Record
- **POST** `/api/v1/records`
  - Creates a new health record
  - Request Body:
  ```json
  {
    "patient_name": "Alice Johnson",
    "age": 42,
    "diagnosis": "High Blood Pressure",
    "treatment": "Medication and lifestyle changes"
  }
  ```
  - Response: `201 Created`

### Update Record
- **PUT** `/api/v1/records/:id`
  - Updates an existing health record
  - Request Body:
  ```json
  {
    "patient_name": "Alice Johnson",
    "age": 42,
    "diagnosis": "Controlled High Blood Pressure",
    "treatment": "Continued medication"
  }
  ```
  - Response: `200 OK` or `404 Not Found`

### Delete Record
- **DELETE** `/api/v1/records/:id`
  - Deletes a health record
  - Response: `200 OK` or `404 Not Found`
  ```json
  {
    "message": "Record deleted successfully"
  }
  ```

## Usage Examples

### Using cURL

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Get all records
curl http://localhost:8080/api/v1/records

# Get a specific record
curl http://localhost:8080/api/v1/records/1

# Create a new record
curl -X POST http://localhost:8080/api/v1/records \
  -H "Content-Type: application/json" \
  -d '{
    "patient_name": "Bob Williams",
    "age": 55,
    "diagnosis": "Diabetes",
    "treatment": "Insulin therapy"
  }'

# Update a record
curl -X PUT http://localhost:8080/api/v1/records/1 \
  -H "Content-Type: application/json" \
  -d '{
    "patient_name": "John Doe",
    "age": 35,
    "diagnosis": "Recovered from Flu",
    "treatment": "No treatment needed"
  }'

# Delete a record
curl -X DELETE http://localhost:8080/api/v1/records/1
```

## Data Model

### HealthRecord

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| id | int | Auto-generated | - |
| patient_name | string | Yes | - |
| age | int | Yes | 0-150 |
| diagnosis | string | Yes | - |
| treatment | string | No | - |

## Project Structure

```
go-healthcare/
├── main.go           # Main application file with API implementation
├── go.mod            # Go module dependencies
├── go.sum            # Dependency checksums
├── .gitignore        # Git ignore file
└── README.md         # This file
```

## Building

To build the application:

```bash
go build -o go-healthcare
./go-healthcare
```

## Best Practices Demonstrated

1. **RESTful API Design**: Proper use of HTTP methods and status codes
2. **Data Validation**: Input validation using struct tags
3. **Error Handling**: Consistent error responses
4. **Code Organization**: Clear separation of concerns
5. **Documentation**: Comprehensive README and code comments
6. **Standard Project Structure**: Following Go conventions

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.