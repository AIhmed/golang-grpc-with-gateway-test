# gRPC Test Project

This is a simple gRPC project with HTTP gateway support that demonstrates how to set up a gRPC service with REST API endpoints.

## Project Structure

```
grpc-test/
├── proto/
│   ├── google/api/
│   │   ├── annotations.proto
│   │   └── http.proto
│   └── testdata/
│       └── testdata.proto
├── server/
│   └── server.go
├── main.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Prerequisites

Make sure you have the following tools installed:

1. **Go** (version 1.24 or later)
2. **Protocol Buffers Compiler (protoc)**
3. **Go protobuf plugins**:
   - `protoc-gen-go`
   - `protoc-gen-go-grpc`
   - `protoc-gen-grpc-gateway`

### Installing protoc and plugins

```bash
# Install protoc (if not already installed)
# On Ubuntu/Debian:
sudo apt-get install protobuf-compiler

# On macOS:
brew install protobuf

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```

## Setup and Usage

1. **Generate protobuf files**:
   ```bash
   make proto
   ```

2. **Run the application**:
   ```bash
   make run
   ```

3. **Or build the application**:
   ```bash
   make build
   ```

## What the Application Does

This application provides:

- **gRPC Server**: Running on port 50051
- **HTTP Gateway**: Running on port 8080, providing REST API endpoints

### API Endpoints

- **gRPC**: `localhost:50051` - TestDataService.GetTestData
- **HTTP**: `GET http://localhost:8080/test-data` - Returns a list of test persons

### Testing the API

You can test the HTTP endpoint using curl:

```bash
curl http://localhost:8080/test-data
```

Expected response:
```json
{
  "persons": [
    {
      "name": "John Doe",
      "age": 30,
      "profession": "Software Engineer"
    },
    {
      "name": "Jane Smith",
      "age": 28,
      "profession": "Data Scientist"
    },
    {
      "name": "Bob Johnson",
      "age": 45,
      "profession": "Product Manager"
    }
  ]
}
```

## Key Changes Made

1. **Fixed import paths**: Updated the protobuf import to use the correct path structure
2. **Added Google API proto files**: Included the necessary `annotations.proto` and `http.proto` files
3. **Updated server code**: Fixed import paths to use the generated testdata package
4. **Added Makefile**: Simplified the build and run process
5. **Added documentation**: Created this README with setup instructions

## Troubleshooting

If you encounter the "google/api/annotations.proto: File not found" error:

1. Make sure you're running `protoc` from the project root directory
2. Verify that the `proto/google/api/` directory contains the required files
3. Check that the import path in `testdata.proto` matches your file structure

The key fix was ensuring that the import path in the protobuf file matches the actual file structure relative to where `protoc` is run from. 