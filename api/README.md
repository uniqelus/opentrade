# OpenTrade Platform API

## Preferences

This section outlines the project preferences, configurations, and required prerequisites for local development and code generation.

### Prerequisites
Ensure the following tools are installed and properly configured on your system:
- **Go**: Version 1.21 or higher.
- **Buf CLI**: The primary tool for managing Protocol Buffers and generating code.
- **Docker**: Required for running the Swagger UI locally (optional, used only for viewing API documentation).

### Project Configuration
The project utilizes `buf` for Protocol Buffers management. The configuration is strictly defined in the following files:
- `buf.yaml`: Defines the module path (`proto`), remote dependencies (`googleapis`, `grpc-gateway`, `protoc-gen-validate`), and linting rules (using the `STANDARD` rule set with specific exceptions for RPC naming).
- `buf.gen.yaml`: Configures the code generation plugins and their respective output directories. Go source code is output to `gen/go`, and OpenAPI specifications are output to `openapi`.

## How to generate files from Protocol Buffers Contracts

To generate the Go source code, gRPC gateway routes, validation logic, and OpenAPI specifications from the `.proto` files, follow these steps:

### 1. Install Required Plugins
The `buf.gen.yaml` configuration uses local plugins. Ensure all necessary Protocol Buffer compiler plugins are installed and available in your system's `$PATH`. Execute the following commands:

```bash
# Core Protobuf and gRPC plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# gRPC-Gateway and OpenAPI plugins
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# Validation plugin
go install github.com/envoyproxy/protoc-gen-validate@latest
```

### 2. Fetch Dependencies
Update the `buf.lock` file to ensure all remote dependencies are synchronized and up to date:

```bash
buf dep update
```

### 3. Generate Code
Run the `buf generate` command from the root of the `api` directory. This will process the `.proto` files located in the `proto/` directory and output the generated files to `gen/go/` and `openapi/`:

```bash
buf generate
```

*Note: The generated files are configured to use `source_relative` paths, meaning they will be placed in directories mirroring the source `.proto` files within the designated output folders.*

## How to launch Swagger Client

The OpenAPI (Swagger) specifications are automatically generated in JSON format within the `openapi/` directory. You can view the interactive API documentation using one of the following methods.

### Method 1: Using Docker (Recommended for Local Development)
The most efficient way to view the Swagger UI without modifying the application code is by using the official Swagger UI Docker image.

Run the following command from the root of the `api` directory:

```bash
docker run -d --name swagger-ui \
  -p 8081:8080 \
  -v $(pwd)/openapi:/usr/share/nginx/html/openapi \
  -e SWAGGER_JSON=/usr/share/nginx/html/openapi/identity_provider/v1/user_service.swagger.json \
  swaggerapi/swagger-ui
```

Once the container is running, access the Swagger UI in your browser at: `http://localhost:8081`.

*(Note: If you are using Windows PowerShell, replace `$(pwd)` with `${PWD}`.)*

### Method 2: Utilizing the go-swagger CLI Tool
For developers who prefer a native Go-based command-line utility over containerized solutions, the go-swagger project provides an efficient mechanism to serve the Swagger UI locally.
#### 1. Installation
Ensure that your Go environment is correctly configured and that the Go binary path (typically $GOPATH/bin or $HOME/go/bin) is included in your system's $PATH. Install the swagger CLI tool by executing the following command:
```bash
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```
Verify the installation by checking the version:
```bash
swagger version
```
#### 2. Serving the Documentation
Navigate to the root directory of the api project and execute the serve command, specifying the desired port and the path to the generated OpenAPI specification:
```bash
swagger serve --port 8081 openapi/identity_provider/v1/user_service.swagger.json
```

The utility will automatically launch the Swagger UI in your default web browser. The documentation will be hosted locally and accessible at http://localhost:8081/docs. To prevent the browser from opening automatically, append the --no-open flag to the command.