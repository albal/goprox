# goprox - HTTP to HTTPS Proxy Server

**goprox** is a lightweight HTTP proxy server written in Go that automatically converts HTTP requests to HTTPS requests. It's designed to help bypass defective SSL/TLS implementations by acting as a local proxy that handles the HTTPS conversion transparently.

## Purpose

Many applications and systems have defective or problematic SSL/TLS implementations that can cause connectivity issues when making HTTPS requests directly. **goprox** solves this by:

- Acting as a local HTTP proxy that you can configure your applications to use
- Automatically converting all incoming HTTP requests to HTTPS requests
- Handling SSL/TLS negotiation on behalf of your applications
- Providing a simple workaround for SSL/TLS implementation problems

## Features

- 🔒 **HTTP to HTTPS Conversion**: Automatically converts HTTP requests to HTTPS
- 🚀 **Lightweight**: Single binary with no external dependencies
- 📝 **Request Logging**: Logs all incoming requests for debugging
- 🔧 **Header Management**: Properly forwards and manages HTTP headers
- ⚡ **High Performance**: Built with Go's efficient HTTP handling

## Installation

### Prerequisites

- Go 1.24.6 or later

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/albal/goprox.git
cd goprox
```

2. Initialize the Go module and build:
```bash
go mod init goprox
go mod tidy
go build -o goprox .
```

3. (Optional) Build with race detection for development:
```bash
go build -race -o goprox-race .
```

## Usage

### Starting the Proxy Server

Run the proxy server:
```bash
./goprox
```

Or run directly with Go:
```bash
go run .
```

The server will start and listen on **port 8888** by default. You should see:
```
Starting HTTP proxy server on :8888
```

### Configuring Your Applications

Configure your applications to use the proxy:

- **Proxy Host**: `localhost` or `127.0.0.1`  
- **Proxy Port**: `8888`
- **Protocol**: HTTP proxy

### Example Usage Scenarios

#### curl
```bash
# Use goprox to make HTTPS requests via HTTP proxy
curl -x http://localhost:8888 http://example.com
```

#### Environment Variables
```bash
export HTTP_PROXY=http://localhost:8888
export HTTPS_PROXY=http://localhost:8888
# Your applications will now use the proxy
```

#### Application Configuration
Many applications support proxy configuration through settings or command-line options. Configure them to use `http://localhost:8888` as the HTTP proxy.

## How It Works

1. **goprox** listens for HTTP requests on port 8888
2. For each incoming request:
   - Extracts the target host and path
   - Converts the request to use HTTPS scheme
   - Forwards the request to the destination server using HTTPS
   - Returns the response back to the client
3. All SSL/TLS negotiation is handled by **goprox**, not your application

### Request Flow
```
Your App → HTTP Request → goprox :8888 → HTTPS Request → Target Server
Your App ← HTTP Response ← goprox ← HTTPS Response ← Target Server
```

## Development

### Code Quality Tools

Format code:
```bash
go fmt ./...
```

Vet code:
```bash
go vet ./...
```

Install and run linter (optional):
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
$HOME/go/bin/golangci-lint run
```

### Project Structure

```
.
├── main.go          # Single source file containing the complete proxy implementation
├── go.mod           # Go module definition
├── LICENSE          # MIT License
├── README.md        # This file
└── .gitignore       # Git ignore rules
```

### Architecture

The proxy is implemented as a single Go file (`main.go`) using only the Go standard library:

- `net/http` - HTTP server and client functionality
- `io` - Request/response body copying
- `log` - Request logging

## Configuration

Currently, **goprox** uses sensible defaults:

- **Port**: 8888 (hardcoded)
- **Protocol**: Converts all requests to HTTPS
- **Logging**: Enabled by default

Future versions may include configuration options for port, logging levels, and protocol handling.

## Troubleshooting

### Port Already in Use
If you see "bind: address already in use", another process is using port 8888:
```bash
# Kill existing goprox processes
pkill goprox
# Or find what's using the port
lsof -i :8888
```

### Connection Issues
- Ensure your application is configured to use HTTP proxy (not HTTPS proxy) 
- Verify the proxy address is `http://localhost:8888`
- Check the proxy logs for error messages

## Contributing

We welcome contributions! Here's how you can help:

### Pull Requests
- 🐛 **Bug fixes** - Help us improve reliability
- ✨ **New features** - Add configuration options, protocol support, etc.
- 📚 **Documentation** - Improve README, add examples
- 🧪 **Tests** - Add test coverage
- 🔧 **Performance** - Optimize proxy performance

### Development Workflow

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes
4. Test your changes: `go build -o goprox . && ./goprox`
5. Format and vet: `go fmt ./... && go vet ./...`
6. Commit your changes: `git commit -am 'Add your feature'`
7. Push to the branch: `git push origin feature/your-feature`
8. Create a Pull Request

### Ideas for Contributions

- Configuration file support
- Command-line flags for port and settings
- HTTPS proxy support (in addition to HTTP proxy)
- Request/response filtering and modification
- Performance metrics and monitoring
- Docker container support
- Authentication support

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Al West** - Initial work

## Acknowledgments

- Built with Go's excellent standard library
- Inspired by the need to work around problematic SSL/TLS implementations
- Thanks to the Go community for excellent documentation and examples