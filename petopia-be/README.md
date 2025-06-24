# Petopia Backend

## Development Flow

This guide provides a comprehensive development flow for setting up and working with the Petopia backend, including Swagger, database setup, migrations, and Docker.

### Prerequisites

- Docker
- Golang (Ensure Go is installed and added to your system's PATH. You can download it from [golang.org](https://golang.org/dl/). After installation, verify by running `go version` in your terminal.)
- golang-migrate (Install using Homebrew: `brew install golang-migrate`)
- swag (Install using Go: `go get -u github.com/swaggo/swag/cmd/swag`)

### Database Setup

1. **Build and Run the PostgreSQL Docker Container**

   Navigate to the `petopia/petopia-be/` directory and build the Docker container:

   ```bash
   docker build -t petopia-db -f Dockerfile.postgresql .
   ```

   Run the Docker container:

   ```bash
   docker run --name petopia-db -d -p 5432:5432 petopia-db
   ```

2. **Configure Environment Variables**

   Update the `.env` file with your database credentials:

   ```plaintext
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_HOST=127.0.0.1:5432
   DB_NAME=petopia
   ```

### Migrations

1. **Apply Migrations**

   Use `golang-migrate` to apply migrations:

   ```bash
   migrate -path ./migrations -database "postgres://user:password@localhost:5432/petopia?sslmode=disable" up
   ```

### Swagger Setup

1. **Generate Swagger Documentation**

   Navigate to the `petopia/petopia-be/` directory and run:

   ```bash
   swag init
   ```

2. **Access Swagger UI**

   Once the application is running, access the Swagger UI at `http://localhost:8080/swagger/index.html`.

### Application Execution

1. **Install GORM and PostgreSQL Driver**

   Navigate to the `petopia/petopia-be/` directory and run:

   ```bash
   go get -u gorm.io/gorm gorm.io/driver/postgres
   ```

2. **Run the Golang Application**

   Navigate to the `petopia/petopia-be/` directory and run the application:

   ```bash
   go run main.go
   ```

   The application will start on port 8080.

3. **Access the API**

   You can access the API at `http://localhost:8080/api`.

### Notes

- Ensure Docker is running on your machine before starting the database setup.
- Replace placeholder values in the `.env` file and Dockerfile with your actual credentials.
- If you encounter issues with the `go` command, ensure Go is installed and properly configured in your system's PATH.
- Use `golang-migrate` to manage database schema versions.
- Use `swag` to generate and update Swagger documentation.


```docker run --name petopia-db -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=P@ssw0rd -e POSTGRES_DB=petopia -p 5432:5432 -d postgres```