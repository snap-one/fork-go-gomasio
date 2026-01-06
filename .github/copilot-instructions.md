# Copilot Instructions for fork-go-gomasio

## Project Overview

Gomasio Fork - A forked Go implementation of Socket.IO and Engine.IO protocols.
This library provides real-time bidirectional communication capabilities used
by OvrC for device and client communication.

**Note**: This is a fork. Changes should be minimal and focused on OvrC-specific
needs. Consider upstreaming general improvements.

## Technology Stack

- **Language**: Go 1.23+
- **Protocols**: Socket.IO, Engine.IO
- **Purpose**: Real-time communication library

## Project Structure

```
├── conn.go              # Connection handling
├── engineio/            # Engine.IO implementation
│   ├── transport/       # Transport layer (WebSocket, polling)
│   └── parser/          # Packet parsing
├── example/             # Usage examples
├── go.mod               # Module definition
└── go.sum               # Dependency checksums
```

## Code Conventions

### Go Standards
- Follow idiomatic Go patterns
- Maintain compatibility with upstream
- Document OvrC-specific changes
- Use interfaces for flexibility

### Fork Maintenance
- Tag OvrC-specific changes with comments
- Keep modifications minimal
- Track upstream version
- Document divergence reasons

### Error Handling
- Return descriptive errors
- Handle connection errors gracefully
- Implement reconnection logic
- Log protocol errors

## API Usage

### Create Connection
```go
conn, err := gomasio.Dial("ws://server:port/socket.io/")
if err != nil {
    return fmt.Errorf("connection failed: %w", err)
}
defer conn.Close()
```

### Send Message
```go
err := conn.Emit("event", data)
```

### Receive Message
```go
conn.On("event", func(data []byte) {
    // Handle incoming event
})
```

## Testing Requirements

### Unit Testing
- Run tests: `go test ./...`
- Test connection handling
- Test protocol compliance
- Test error scenarios

### Integration Testing
- Test with real Socket.IO server
- Test reconnection behavior
- Test transport fallback

## Development Workflow

### Local Setup
1. Ensure Go 1.23+ installed
2. Run: `go mod download`
3. Run tests: `go test ./...`

### Build
```bash
go build ./...
```

## Engine.IO Transports

Supported transports:
- WebSocket (primary)
- HTTP long-polling (fallback)
- Both with proper upgrade handling

## Socket.IO Features

- Event-based messaging
- Acknowledgements
- Namespaces
- Room support

## Security Considerations

- Validate server certificates
- Use secure WebSocket (wss://)
- Handle authentication properly
- Sanitize event data
- Implement rate limiting
