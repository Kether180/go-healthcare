# Testing the Go Health API

This guide shows you how to test all the API endpoints using different methods.

## Method 1: Using curl (Command Line)

### 1. Health Check
```bash
curl http://localhost:8080/health
```

### 2. Get All Users
```bash
curl http://localhost:8080/api/v1/users
```

### 3. Get User by ID
```bash
curl http://localhost:8080/api/v1/users/1
```

### 4. Create New User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "David Wilson",
    "email": "david@example.com"
  }'
```

### 5. Update User
```bash
curl -X PUT http://localhost:8080/api/v1/users/4 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "David Wilson Updated",
    "email": "david.updated@example.com"
  }'
```

### 6. Delete User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/4
```

---

## Method 2: Using PowerShell (Windows)

### 1. Health Check
```powershell
Invoke-RestMethod -Uri http://localhost:8080/health -Method Get
```

### 2. Get All Users
```powershell
Invoke-RestMethod -Uri http://localhost:8080/api/v1/users -Method Get
```

### 3. Get User by ID
```powershell
Invoke-RestMethod -Uri http://localhost:8080/api/v1/users/1 -Method Get
```

### 4. Create New User
```powershell
$body = @{
    name = "David Wilson"
    email = "david@example.com"
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/v1/users `
  -Method Post `
  -Body $body `
  -ContentType "application/json"
```

### 5. Update User
```powershell
$body = @{
    name = "David Wilson Updated"
    email = "david.updated@example.com"
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/v1/users/4 `
  -Method Put `
  -Body $body `
  -ContentType "application/json"
```

### 6. Delete User
```powershell
Invoke-RestMethod -Uri http://localhost:8080/api/v1/users/4 -Method Delete
```

---

## Method 3: Using VS Code REST Client Extension

Install the "REST Client" extension in VS Code, then create a file called `api-tests.http`:

```http
### Health Check
GET http://localhost:8080/health

### Get All Users
GET http://localhost:8080/api/v1/users

### Get User by ID
GET http://localhost:8080/api/v1/users/1

### Create New User
POST http://localhost:8080/api/v1/users
Content-Type: application/json

{
  "name": "David Wilson",
  "email": "david@example.com"
}

### Update User
PUT http://localhost:8080/api/v1/users/4
Content-Type: application/json

{
  "name": "David Wilson Updated",
  "email": "david.updated@example.com"
}

### Delete User
DELETE http://localhost:8080/api/v1/users/4
```

Click "Send Request" above each request to execute it.

---

## Method 4: Using a Browser (for GET requests only)

Simply open these URLs in your browser:

- Health Check: http://localhost:8080/health
- Get All Users: http://localhost:8080/api/v1/users
- Get User by ID: http://localhost:8080/api/v1/users/1

---

## Expected Responses

### Health Check
```json
{
  "status": "healthy",
  "message": "Service is running",
  "version": "1.0.0"
}
```

### Get All Users
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

### Get User by ID
```json
{
  "data": {
    "id": 1,
    "name": "Alice Johnson",
    "email": "alice@example.com"
  }
}
```

### Create User (Success)
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

### Update User (Success)
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

### Delete User (Success)
```json
{
  "message": "User deleted successfully"
}
```

---

## Error Responses

### Invalid User ID
```json
{
  "error": "Invalid user ID format"
}
```

### User Not Found
```json
{
  "error": "User not found"
}
```

### Invalid Request Body
```json
{
  "error": "Invalid request body"
}
```

### Missing Required Fields
```json
{
  "error": "Name and email are required"
}
```

---

## Testing Workflow

1. **Start the server**: `go run main.go`
2. **Test health check**: Verify server is running
3. **Get all users**: See initial data
4. **Create a new user**: Add user ID 4
5. **Get user by ID**: Verify creation
6. **Update the user**: Modify user 4
7. **Delete the user**: Remove user 4
8. **Try to get deleted user**: Should return 404

---

## Tips for Testing

- Use `jq` to format JSON responses: `curl http://localhost:8080/health | jq`
- Use `-v` flag with curl to see full HTTP request/response: `curl -v http://localhost:8080/health`
- Watch server logs in the terminal to see requests being processed
- Test error cases (invalid IDs, missing fields) to ensure proper error handling
