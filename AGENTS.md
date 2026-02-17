# Agent Instructions

## Build, Test, and Run Commands

### Running Locally

```bash
# Copy environment configuration
cp .env_backup .env
source .env

# Start dependencies (MySQL, Redis)
make run-dependencies

# Run migrations (from app directory)
cd app && make run-migration

# Run service with live reload (from app directory)
cd app && make run-service  # Uses gin for hot-reload on port 8082
```

### Docker

```bash
# Build Docker image
make build

# Run entire stack (dependencies + migration + service)
make run-service

# Run integration tests
make run-integration-testing
```

### Database Migrations

```bash
# Create new migration
make create-migration  # Prompts for migration name

# Migrations are in ./migrations directory
# Apply via: app/main.go with args: false service.log run-migration
```

## Architecture

This is a **Go monolith** using **Clean Architecture** with domain-driven organization.

### Domain Structure

All business domains live in `handler/` directory. Each domain (e.g., `security`, `users`, `template`) contains:

- **delivery/** - Input/output handling (REST, gRPC). Defines endpoints and request/response marshalling
- **usecase/** - Business logic layer. Implements domain interfaces
- **entity/** - Repository layer. Database access via GORM
- **models/** - Database models (structs matching DB schema)
- **payload/** - API request/response DTOs
- **interface.go** - Domain interfaces (e.g., `IUserUseCase`, `IUserRepo`)

### Dependency Flow

```text
delivery -> usecase -> entity/repo -> database
```

### Service Initialization

1. `app/main.go` - Entry point, sets up logging, config, tracing
2. `app/service.go` - Wire up dependencies:
   - Create database/Redis clients
   - Initialize entities (repos)
   - Initialize usecases
   - Register delivery endpoints

### Key Infrastructure

- **Database**: MySQL 8.0.33 via GORM (with tracing via `otelgorm`)
- **Cache**: Redis v8 (with tracing via `redisotel`)
- **HTTP Framework**: Gin with middleware for CORS, request ID, OpenTelemetry
- **Migrations**: golang-migrate with SQL files
- **Observability**: Uptrace for distributed tracing
- **Live Reload**: codegangsta/gin (runs app on port 8082, proxy on 8080)

## Key Conventions

### Configuration Management

- Config defined in `config/configuration.go` with struct tags for JSON and ENV
- Uses hybrid approach: JSON file (`config/local.conf`) + environment variable overrides
- Load via `github.com/muhammadrivaldy/go-util`

### Error Handling

- Custom error type: `util.Error` with fields `Error`, `Code`, `Object`
- Return pattern: `(result, util.Error)` instead of standard `(result, error)`
- HTTP responses via `goutil.ResponseError(c, code, err, object)`

### Context Management

- Custom context utility: `util.GetContext(ctx)` extracts user info from request context
- User info includes `UserID` parsed from JWT

### API Registration Pattern

Security domain has an API registration system:

- Each endpoint has a UUID and is registered in database
- `useCaseSecurity.RegisterService()` registers service on startup
- Endpoints defined as `payload.RequestRegisterApi` structs with ID, Name, Endpoint, Method, ServiceID

### Middleware

- `middleware.ValidateAccess(apiID)` - JWT auth + RBAC check per API endpoint
- Applied to protected routes via Gin middleware chain

### Logging

- Uses `go.uber.org/zap`
- Configured in `logs/` package
- Output controlled by command-line args to `app/main.go`

### Module Path

Import path is `backend`, not `go-monolith` (see go.mod line 1)

### Adding New Domains

1. Create `handler/<domain>/` directory
2. Add subdirs: `delivery/`, `usecase/`, `entity/`, `models/`, `payload/`
3. Define interfaces in `handler/<domain>/interface.go`
4. Wire up in `app/service.go` following existing pattern
