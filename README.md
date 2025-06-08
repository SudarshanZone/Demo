# Complete Swagger Annotations Guide

## 1. Main Application File (main.go)

### Basic API Information
```go
// @title Your API Name
// @version 1.0
// @description This is a sample server for your API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support  
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
```

### Security Definitions
```go
// Bearer Token Authentication
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345"

// Basic Authentication
// @securityDefinitions.basic BasicAuth

// OAuth2 Authentication
// @securityDefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// API Key Authentication
// @securityDefinitions.apikey ApiKeyAuth
// @in query
// @name api_key
```

---

## 2. Handler/Controller Files

### Function-Level Annotations

#### Basic Handler Annotations
```go
// @Summary Short description of the endpoint
// @Description Detailed description of what this endpoint does
// @Tags tag1,tag2
// @ID uniqueOperationId
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /users [get]
// @Success 200 {object} ResponseModel
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
```

#### Parameters
```go
// Path Parameters
// @Param id path int true "User ID"

// Query Parameters  
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param status query string false "Status filter" Enums(active, inactive)

// Header Parameters
// @Param Authorization header string true "Bearer token"
// @Param Content-Type header string true "Content type" default(application/json)

// Body Parameters
// @Param user body UserCreateRequest true "User data"

// Form Parameters
// @Param name formData string true "User name"
// @Param email formData string true "User email"
// @Param avatar formData file false "User avatar"
```

#### Response Examples
```go
// @Success 200 {object} User "Successfully retrieved user"
// @Success 201 {object} User "User created successfully"  
// @Success 204 "No content"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 404 {object} ErrorResponse "Not found"
// @Failure 422 {object} ValidationError "Validation error"
// @Failure 500 {object} ErrorResponse "Internal server error"
```

---

## 3. Model/Struct Files

### Struct Tags
```go
type User struct {
    ID        uint      `json:"id" example:"1" format:"int64"`
    Name      string    `json:"name" example:"John Doe" validate:"required,min=2,max=100"`
    Email     string    `json:"email" example:"john@example.com" format:"email" validate:"required,email"`
    Age       int       `json:"age" example:"25" minimum:"0" maximum:"150"`
    Status    string    `json:"status" example:"active" enums:"active,inactive"`
    CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z" format:"date-time"`
    UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z" format:"date-time"`
    Profile   *Profile  `json:"profile,omitempty"`
    Tags      []string  `json:"tags" example:"tag1,tag2"`
    Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type Profile struct {
    Bio       string `json:"bio" example:"Software developer"`
    Website   string `json:"website" example:"https://example.com" format:"uri"`
    Location  string `json:"location" example:"New York"`
    AvatarURL string `json:"avatar_url" example:"https://example.com/avatar.jpg" format:"uri"`
}
```

### Request/Response Models
```go
type UserCreateRequest struct {
    Name     string `json:"name" binding:"required,min=2,max=100" example:"John Doe"`
    Email    string `json:"email" binding:"required,email" example:"john@example.com"`
    Password string `json:"password" binding:"required,min=8" example:"password123"`
    Age      int    `json:"age" binding:"min=0,max=150" example:"25"`
}

type UserUpdateRequest struct {
    Name  *string `json:"name,omitempty" example:"John Doe"`
    Email *string `json:"email,omitempty" example:"john@example.com"`
    Age   *int    `json:"age,omitempty" example:"25"`
}

type UserResponse struct {
    ID        uint      `json:"id" example:"1"`
    Name      string    `json:"name" example:"John Doe"`
    Email     string    `json:"email" example:"john@example.com"`
    Age       int       `json:"age" example:"25"`
    Status    string    `json:"status" example:"active"`
    CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
}

type ListResponse struct {
    Data       []UserResponse `json:"data"`
    Pagination Pagination     `json:"pagination"`
}

type Pagination struct {
    Page       int `json:"page" example:"1"`
    Limit      int `json:"limit" example:"10"`
    Total      int `json:"total" example:"100"`
    TotalPages int `json:"total_pages" example:"10"`
}

type ErrorResponse struct {
    Error   string `json:"error" example:"Something went wrong"`
    Message string `json:"message" example:"Detailed error message"`
    Code    int    `json:"code" example:"400"`
}
```

---

## 4. Complete Handler Examples

### GET Endpoint
```go
// GetUsers retrieves a list of users
// @Summary Get all users
// @Description Get a paginated list of all users with optional filtering
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page number" default(1) minimum(1)
// @Param limit query int false "Items per page" default(10) minimum(1) maximum(100)
// @Param status query string false "Filter by status" Enums(active, inactive)
// @Param search query string false "Search by name or email"
// @Success 200 {object} ListResponse "Successfully retrieved users"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users [get]
func GetUsers(c *gin.Context) {
    // Handler implementation
}
```

### POST Endpoint
```go
// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided data
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body UserCreateRequest true "User creation data"
// @Success 201 {object} UserResponse "User created successfully"
// @Failure 400 {object} ErrorResponse "Bad request - validation error"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 409 {object} ErrorResponse "Conflict - user already exists"
// @Failure 422 {object} ValidationError "Validation error"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users [post]
func CreateUser(c *gin.Context) {
    // Handler implementation
}
```

### PUT Endpoint
```go
// UpdateUser updates an existing user
// @Summary Update user by ID
// @Description Update user information by user ID
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID" format(int64) minimum(1)
// @Param user body UserUpdateRequest true "User update data"
// @Success 200 {object} UserResponse "User updated successfully"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 422 {object} ValidationError "Validation error"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
    // Handler implementation
}
```

### DELETE Endpoint
```go
// DeleteUser deletes a user
// @Summary Delete user by ID
// @Description Delete a user by their ID
// @Tags users
// @Security ApiKeyAuth
// @Param id path int true "User ID" format(int64) minimum(1)
// @Success 204 "User deleted successfully"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
    // Handler implementation
}
```

### File Upload Endpoint
```go
// UploadUserAvatar uploads user avatar
// @Summary Upload user avatar
// @Description Upload an avatar image for a user
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param avatar formData file true "Avatar image file"
// @Param description formData string false "Image description"
// @Success 200 {object} UserResponse "Avatar uploaded successfully"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 413 {object} ErrorResponse "File too large"
// @Failure 415 {object} ErrorResponse "Unsupported media type"
// @Router /users/{id}/avatar [post]
func UploadUserAvatar(c *gin.Context) {
    // Handler implementation
}
```

---

## 5. Advanced Annotations

### Custom Response Types
```go
type APIResponse struct {
    Success bool        `json:"success" example:"true"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty" example:""`
    Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
    RequestID string    `json:"request_id" example:"req-123"`
    Timestamp time.Time `json:"timestamp" example:"2023-01-01T00:00:00Z"`
    Version   string    `json:"version" example:"1.0"`
}
```

### Nested Objects
```go
type UserWithProfile struct {
    User
    Profile Profile `json:"profile"`
    Roles   []Role  `json:"roles"`
}

type Role struct {
    ID          uint   `json:"id" example:"1"`
    Name        string `json:"name" example:"admin"`
    Description string `json:"description" example:"Administrator role"`
}
```

### Enums and Validation
```go
type OrderStatus string

const (
    OrderStatusPending   OrderStatus = "pending"
    OrderStatusConfirmed OrderStatus = "confirmed"
    OrderStatusShipped   OrderStatus = "shipped"
    OrderStatusDelivered OrderStatus = "delivered"
    OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
    ID     uint        `json:"id" example:"1"`
    Status OrderStatus `json:"status" example:"pending" enums:"pending,confirmed,shipped,delivered,cancelled"`
    Amount float64     `json:"amount" example:"99.99" minimum:"0"`
}
```

---

## 6. Generation Commands

```bash
# Initialize Swagger in your project
swag init

# Generate with custom output directory
swag init -o ./docs

# Generate with custom main file
swag init -g cmd/main.go

# Generate with custom exclude directories
swag init --exclude ./vendor,./tests

# Generate with custom parse dependency
swag init --parseDependency --parseInternal

# Generate with specific output format
swag init --outputTypes go,json,yaml
```

---

## 7. Common Tags Reference

| Tag | Description | Example |
|-----|-------------|---------|
| `@Summary` | Brief endpoint description | `@Summary Get user by ID` |
| `@Description` | Detailed description | `@Description Retrieve user information by user ID` |
| `@Tags` | Group endpoints | `@Tags users,admin` |
| `@Accept` | Request content type | `@Accept json` |
| `@Produce` | Response content type | `@Produce json` |
| `@Param` | Parameter definition | `@Param id path int true "User ID"` |
| `@Success` | Success response | `@Success 200 {object} User` |
| `@Failure` | Error response | `@Failure 404 {object} ErrorResponse` |
| `@Router` | Route definition | `@Router /users/{id} [get]` |
| `@Security` | Authentication requirement | `@Security ApiKeyAuth` |
| `@ID` | Unique operation ID | `@ID getUserById` |

---

## 8. File Organization Best Practices

```
project/
├── docs/           # Generated swagger files
├── handlers/       # API handlers with swagger annotations
├── models/         # Data models with struct tags
├── middleware/     # Middleware (auth, etc.)
├── main.go        # Main file with API info annotations
└── swagger.yaml   # Generated swagger spec
```
