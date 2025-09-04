# goprox - HTTP Proxy Server

goprox is a simple HTTP proxy server written in Go. It listens on port 8888 and forwards HTTP requests to their intended destinations, handling headers and response proxying.

Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.

## Working Effectively

### Initial Setup (REQUIRED FIRST STEPS)
- Initialize Go module: `go mod init goprox` (CRITICAL: Must be run first - the repository doesn't include go.mod)
- Tidy dependencies: `go mod tidy`
- Verify Go version: `go version` (tested with Go 1.24.6)

### Building and Testing
- Build the application: `go build -o goprox .` -- takes ~1 second (first build may take 10+ seconds). Set timeout to 30+ seconds.
- Build with race detection: `go build -race -o goprox-race .` -- takes ~1 second. Set timeout to 30+ seconds.
- Run directly: `go run .` -- starts server immediately, no build artifacts
- Format code: `go fmt ./...` -- instant
- Vet code: `go vet ./...` -- takes ~2 seconds
- Run tests: `go test -v .` -- takes ~2 seconds (reports "no test files" - this is expected)

### Linting and Code Quality
- Install golangci-lint: `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest` -- takes ~30 seconds. Set timeout to 60+ seconds.
- Run linter: `$HOME/go/bin/golangci-lint run` -- takes ~3 seconds, finds 2 errcheck issues (non-blocking)
- Always run linting before committing - the current code has minor issues but builds successfully

### Running the Application
- Start server: `./goprox` or `go run .`
- Server listens on port 8888
- Logs all incoming requests with method, host, and path
- To stop: Ctrl+C or kill the process

## Validation

### CRITICAL: Build and Run Validation
After making ANY changes to the code:
1. ALWAYS run the complete setup: `go mod init goprox && go mod tidy` (if not done)
2. ALWAYS build: `go build -o goprox .` and verify it succeeds
3. ALWAYS test server starts: `./goprox &` then verify logs show "Starting HTTP proxy server on :8888"
4. ALWAYS stop the server: `kill %1` or Ctrl+C
5. NEVER skip these steps - they are essential for validating the proxy works

### Manual Testing Scenarios
- **Basic functionality**: Start server, confirm it logs "Starting HTTP proxy server on :8888"
- **Request handling**: Server should log received requests (you can see this in output)
- **Clean shutdown**: Server should stop cleanly with Ctrl+C

### Code Quality Checks
- Always run `go fmt ./...` before committing
- Always run `go vet ./...` before committing  
- Run `$HOME/go/bin/golangci-lint run` to check for issues (current code has 2 errcheck warnings - acceptable but good to be aware of)

## Build and Timing Information

### Expected Command Times
- `go mod init goprox`: instant
- `go mod tidy`: ~1 second
- `go build -o goprox .`: ~1 second (first build 10+ seconds) - set timeout 30+ seconds
- `go build -race -o goprox-race .`: ~1 second - set timeout 30+ seconds
- `go fmt ./...`: instant
- `go vet ./...`: instant
- `go test -v .`: ~1 second
- `golangci-lint run`: ~1 second
- Server startup: instant

### Binary Information
- Output binary: ~8.9MB
- No external dependencies (uses only Go standard library)
- Builds clean on Linux x86_64

## Repository Structure

### File Overview
```
.
├── .git/           # Git repository
├── .gitignore      # Excludes binaries and common Go artifacts  
├── LICENSE         # MIT License
├── main.go         # Single source file - complete HTTP proxy implementation
└── go.mod          # Go module file (created by go mod init)
```

### Key Code Components (main.go)
- `main()`: Starts HTTP server on :8888 with proxyHandler
- `proxyHandler()`: Core proxy logic - handles request forwarding, header copying, response proxying
- Uses standard library: net/http, io, log
- No external dependencies

## Common Issues and Solutions

### Missing go.mod
**Issue**: Build fails with module errors
**Solution**: Run `go mod init goprox` first

### Port 8888 in use
**Issue**: Server fails to start - "bind: address already in use"
**Solution**: Kill existing process: `pkill goprox` or use different port

### Linting Issues
**Issue**: golangci-lint reports errcheck warnings
**Solution**: These are minor - `resp.Body.Close()` and `io.Copy()` error handling. Non-blocking for builds.

## Development Workflow

For any code changes:
1. Make your changes to main.go
2. Run `go fmt ./...`
3. Run `go vet ./...` 
4. Run `go build -o goprox .` -- set timeout 30+ seconds, builds in ~1 second
5. Test: `./goprox` and verify startup message
6. Stop server and proceed with commit
7. Optionally run `$HOME/go/bin/golangci-lint run` for code quality check

This is a minimal, single-file Go application. Most changes will be in main.go. Always validate the proxy starts correctly after changes.