# Clean Architecture Implementation for Petopia Backend

This document outlines the clean architecture implementation in the Petopia backend, explaining how the different components interact and providing guidelines for maintaining and extending the codebase.

## Overview

Clean Architecture separates concerns into concentric layers with dependencies pointing inward. This ensures:
- Independence from frameworks and external agencies
- Testability of business rules without UI, database, web server, or any external element
- Independence from UI, allowing it to change without changing the system
- Independence from database, allowing the database to be swapped out with minimal effort
- Independence from any external agency, making our business rules not bound to external interfaces

## Architecture Layers

### 1. Domain Models (Core)

**Location:** `/models/mongo`

These are the core business entities representing the concepts and relationships in our application domain. They should be framework-independent and contain only business logic.

Key files:
- `product_details.go`
- `product_brand.go`
- `customer_review.go`

### 2. Repository Layer

**Location:** `/repository`

Abstracts the data access mechanisms from the business logic through interfaces, allowing flexibility in data source implementation.

Key components:
- **Interfaces:** Defined in `interfaces.go`
- **Implementations:** MongoDB-specific implementations in files like `mongo_product_repository.go`

Example interface:
```go
type ProductRepository interface {
    FindAll() ([]models.ProductDetails, error)
    FindById(id primitive.ObjectID) (*models.ProductDetails, error)
    Create(product *models.ProductDetails) (*models.ProductDetails, error)
    Update(id primitive.ObjectID, product *models.ProductDetails) (*models.ProductDetails, error)
    Delete(id primitive.ObjectID) error
    // Other methods
}
```

### 3. Service Layer

**Location:** `/service`

Contains application-specific business rules and orchestrates the flow of data to and from the entities while enforcing business rules.

Key components:
- **Interfaces:** Defined in `service.go` or specific files like `product_service_interface.go`
- **Implementations:** Concrete implementations in files like `service_impl.go` or `product_service.go`

Example interface:
```go
type ProductService interface {
    GetAllProducts() ([]dto.ProductResponse, error)
    GetProductById(id string) (*dto.ProductResponse, error)
    CreateProduct(product dto.ProductRequest) (*dto.ProductResponse, error)
    UpdateProduct(id string, product dto.ProductRequest) (*dto.ProductResponse, error)
    DeleteProduct(id string) error
    // Other methods
}
```

### 4. Controller Layer

**Location:** `/controller`

Handles HTTP requests and responses. Controllers use services to perform business operations and transform the results into HTTP responses.

Key files:
- `product_controller_final.go`
- `message_controller.go`

Example controller method:
```go
func (c *ProductControllerV2) GetAllProducts(w http.ResponseWriter, r *http.Request) {
    products, err := c.service.GetAllProducts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}
```

### 5. Data Transfer Objects (DTOs)

**Location:** `/dto`

Objects used to transfer data between layers, especially between the controller and service layers.

Key files:
- `product_request_dto.go`
- `product_response_dto.go`
- `dto_mapper.go`

Example DTO:
```go
type ProductRequest struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    // Other fields
}
```

### 6. Database Layer

**Location:** `/db`

Manages database connections and provides interfaces for database access.

Key files:
- `db.go`
- `mongodb.go`

## Flow of Control

1. HTTP request comes into a controller
2. Controller parses request and converts to DTO
3. Controller calls appropriate service method with DTO
4. Service performs business logic and calls repository methods
5. Repository interacts with the database using domain models
6. Repository returns domain models to service
7. Service maps domain models to DTOs and returns to controller
8. Controller formats response and sends back to client

## Best Practices

1. **Dependency Injection:** Use dependency injection to provide repositories to services and services to controllers.

2. **Error Handling:** Use consistent error handling throughout the application.

3. **Validation:** Validate input at the controller level before passing to services.

4. **Mapping:** Use mapper functions to convert between DTOs and domain models.

5. **Testing:** Write tests for each layer in isolation using mocks.

## Adding New Features

When adding new features:

1. Define the domain model in `/models/mongo`
2. Create repository interface in `/repository/interfaces.go`
3. Implement the repository in `/repository`
4. Define DTOs in `/dto`
5. Create service interface and implementation
6. Implement controller methods
7. Add routes in `/routes/routes.go`

## Example Implementation

### Domain Model
```go
type Product struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Name        string             `bson:"name"`
    Description string             `bson:"description"`
    Price       float64            `bson:"price"`
    // Other fields
}
```

### Repository Interface
```go
type ProductRepository interface {
    FindAll() ([]models.Product, error)
    FindById(id primitive.ObjectID) (*models.Product, error)
    // Other methods
}
```

### Service Implementation
```go
type productService struct {
    repo repository.ProductRepository
}

func (s *productService) GetAllProducts() ([]dto.ProductResponse, error) {
    products, err := s.repo.FindAll()
    if err != nil {
        return nil, err
    }
    
    return dto.MapProductsToResponseDTOs(products), nil
}
```

### Controller Implementation
```go
type ProductController struct {
    service service.ProductService
}

func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
    products, err := c.service.GetAllProducts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}
```
