.PHONY: run build clean test deps tidy

# Default target
run:
	go run cmd/api/main.go

# Build the API server
build:
	go build -o bin/api cmd/api/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Run tests
test:
	go test ./...

# Download dependencies
deps:
	go mod download

# Tidy dependencies
tidy:
	go mod tidy

# Install new dependencies (from README)
init-deps:
	go get -u github.com/gin-gonic/gin
	go get gopkg.in/yaml.v3
	go get -u go.uber.org/zap
	go get github.com/redis/go-redis/v9