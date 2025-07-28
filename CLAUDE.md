# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Commands

### Build and Run
```bash
# Run API server
go run cmd/api/main.go

# Build API server
go build -o bin/api cmd/api/main.go
```

### Dependencies
```bash
go mod download
go mod tidy
```

### Testing
```bash
go test ./...
```

### Initialize new dependencies (from README)
```bash
go get -u github.com/gin-gonic/gin
go get gopkg.in/yaml.v3
go get -u go.uber.org/zap
go get github.com/redis/go-redis/v9
```

## Project Architecture

This is a Go-based badminton matching service using:
- **Gin**: HTTP web framework for REST API
- **Redis**: Caching layer (configured via config.yml)
- **Zap**: Structured logging
- **YAML**: Configuration management

### Directory Structure
- `cmd/api/main.go`: API server entry point with Gin router setup and basic REST endpoints
- `cmd/`: Directory for multiple application entrypoints
- `configs/config.yml`: Application configuration file
- `config/`: Configuration management package
- `common/cache/`: Redis client wrapper
- `common/logger/`: Logging utilities (currently minimal)
- `domain/player/`: Player domain models (currently empty)
- `domain/match/`: Match domain models (currently empty)

### Configuration
- Configuration is loaded from `configs/config.yml`
- Redis connection string configured in config file
- Development logger used by default

### Current API Endpoints
- `GET /ping`: Health check endpoint
- `GET /user/:name`: Retrieve user data from in-memory map
- `POST /admin`: Create/update user data (requires basic auth: foo/bar or manu/123)

### Development Notes
- The service is in early development stage
- Redis integration is stubbed (setupRouter accepts cache but doesn't use it yet)
- Domain models (player, match) are empty placeholder packages
- Basic auth credentials are hardcoded for development