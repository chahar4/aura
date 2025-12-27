# AGENTS.md

This document outlines the project structure, essential commands, and conventions for working effectively in this codebase.

## Project Structure

The project follows a clean architecture pattern with clear separation of concerns:

-   **`main.go`**: Entry point of the application, responsible for initializing the database, repositories, services, and handlers, and setting up the HTTP router.
-   **`core/`**: Contains the core business logic.
    -   **`core/domains/`**: Defines the application's entities (e.g., `User`, `Guild`) and interfaces for repositories (e.g., `UserRepository`, `GuildRepository`). This layer is framework-agnostic.
    -   **`core/services/`**: Implements the business logic using the domain interfaces (e.g., `UserService`, `GuildService`). These services interact with the repositories defined in the domain layer.
    -   **`core/tools/`**: Contains utility functions and helper modules that can be used across different layers (e.g., `email.go`, `errors.go`, `json.go`).
-   **`adapter/`**: Handles external concerns and adaptations.
    -   **`adapter/handlers/`**: Implements HTTP request handlers (e.g., `UserHandler`, `GuildHandler`) that interact with the core services. These are responsible for parsing requests and formatting responses.
    -   **`adapter/storages/`**: Provides concrete implementations of the repository interfaces defined in `core/domains/`, typically interacting with a database (e.g., `UserPostgresRepo`, `GuildPostgresRepo`).

## Dependencies

The project uses Go modules for dependency management. Key dependencies include:

-   **`github.com/go-chi/chi/v5`**: For HTTP routing.
-   **`gorm.io/driver/postgres` & `gorm.io/gorm`**: For database interaction with PostgreSQL.
-   **`github.com/golang-jwt/jwt/v5`**: For JWT-based authentication.
-   **`github.com/joho/godotenv`**: For loading environment variables from `.env` files.
-   **`golang.org/x/crypto`**: For cryptographic operations, likely password hashing.
-   **`gopkg.in/gomail.v2`**: For sending emails.

## Essential Commands

The `Makefile` provides several commands for development and database management:

-   `make dockerinit`: Initializes a PostgreSQL Docker container.
-   `make createdb`: Creates the `auraDB` database in the PostgreSQL container.
-   `make postgres`: Connects to the PostgreSQL database via `psql`.
-   `make dropdb`: Drops the `auraDB` database.
-   `make migratecreate`: Creates a new migration file.
-   `make migrateup`: Applies pending database migrations.
-   `make migratedown`: Reverts the last database migration.

## Database Migrations

Database migrations are managed using the `migrate` tool. Migration files are located in the `db/migrations` directory.

-   To create a new migration: `make migratecreate name_of_migration_here`
-   To apply migrations: `make migrateup`
-   To revert migrations: `make migratedown`

## Running the Project

1.  Ensure a PostgreSQL database is running and accessible (e.g., using `make dockerinit` and `make createdb`).
2.  Set up environment variables in a `.env` file (e.g., database connection details, JWT secret, email credentials).
3.  Run the application using `go run main.go`. The server will listen on port `3000`.

## Testing

*(No explicit testing framework or patterns were immediately evident in the initial scan. Assume standard Go testing conventions where tests reside in `*_test.go` files within their respective packages.)*
