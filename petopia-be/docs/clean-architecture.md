# Clean Architecture Implementation for Petopia

This document outlines the changes made to implement a clean architecture pattern for the Petopia backend.

## Architecture Overview

The application now follows a strict clean architecture pattern with the following layers:

1. **Controllers** - Handle HTTP requests and responses
2. **Services** - Contain business logic
3. **Repositories** - Provide data access abstraction
4. **Models** - Define domain entities
5. **DTOs** - Handle data transfer between layers

## Key Changes

### 1. Repository Layer

- Created interface definitions in `interfaces.go`
- Implemented MongoDB repositories:
  - `MongoProductRepository`
  - `MongoBrandRepository`
- Added methods for all CRUD operations plus specialized queries
- Added `GetCollection()` method for DAO access

### 2. Service Layer

- Created service interfaces in `service.go`
- Implemented concrete services in `service_impl.go`
- Added `ServiceContainer` for dependency management
- Services now use repositories instead of direct DB access

### 3. DTO Layer

- Enhanced DTOs for request and response handling
- Added mapping functions in `dto_mapper.go`
- Created proper pagination support

### 4. Controller Layer

- Updated controllers to depend on services only
- Removed direct repository access
- Added proper error handling with status codes

## Benefits

1. **Separation of Concerns** - Each layer has a specific responsibility
2. **Testability** - Interfaces allow for easy mocking during tests
3. **Maintainability** - Changes in one layer don't affect others
4. **Flexibility** - Database implementations can be swapped without affecting business logic

## File Structure

```
petopia-be/
├── controller/
│   ├── product_controller.go         # Original controller
│   └── product_controller_v2.go      # New controller using service layer
├── service/
│   ├── service.go                    # Service interfaces
│   ├── product_service.go            # Original service
│   └── service_impl.go               # New service implementations
├── repository/
│   ├── interfaces.go                 # Repository interfaces
│   ├── mongo_product_repository.go   # MongoDB product implementation
│   └── mongo_brand_repository.go     # MongoDB brand implementation
├── dto/
│   ├── product_request_dto.go        # Request DTOs
│   ├── product_response_dto.go       # Response DTOs
│   └── dto_mapper.go                 # DTO <-> Model mappers
└── models/
    └── mongo/
        ├── product_details.go        # Product domain model
        └── product_brand.go          # Brand domain model
```

## Next Steps

1. Complete the transition from DAO to Repository pattern
2. Create unit tests for each layer
3. Implement additional repositories for other entities
4. Update the Swagger documentation

## Example Flow

```
HTTP Request
    ↓
Controller (product_controller_v2.go)
    ↓
Service (service_impl.go)
    ↓
Repository (mongo_product_repository.go)
    ↓
Database
    ↓
Repository (converts DB result to model)
    ↓
Service (applies business logic)
    ↓
Controller (converts to DTO response)
    ↓
HTTP Response
```

This refactoring ensures all database operations go through the repository layer, and all business logic is contained in the service layer, providing a clean separation of concerns.
