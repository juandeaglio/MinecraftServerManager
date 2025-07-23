# Minecraft Server Manager

A robust, production-ready REST API service built in Go for remotely managing Minecraft server processes. This service provides centralized control over Minecraft server lifecycle management with comprehensive monitoring capabilities.

## Overview

The Minecraft Server Manager enables administrators to remotely control Minecraft servers through a simple REST API interface. Whether you're managing a single server or multiple instances, this service provides the infrastructure needed to start, stop, monitor, and query Minecraft servers without direct server access.

## Features

- **Remote Management**: Control Minecraft servers from anywhere without SSH or RDP access
- **Automation Ready**: Perfect for CI/CD pipelines, scheduled maintenance, and automated scaling
- **Monitoring Integration**: Built-in health checks and status monitoring for server uptime tracking
- **Multi-Environment Support**: Suitable for development, staging, and production server management
- **Cost Optimization**: Start/stop servers on-demand to reduce hosting costs during off-peak hours
- **Centralized Control**: Manage multiple Minecraft instances through a single API endpoint


### Core Functionality
- **Process Management**: Start, stop, and monitor Minecraft server processes
- **Health Monitoring**: Real-time server status checking and process validation
- **RCON Integration**: Query server statistics, player counts, and execute server commands
- **REST API**: Clean HTTP endpoints for all server operations
- **Cross-Platform**: Support for Windows and Linux environments

### API Endpoints
- `POST /start` - Start the Minecraft server process
- `POST /stop` - Gracefully stop the server
- `GET /status` - Get detailed server status with player information
- `GET /running` - Quick health check for server process status

### Technical Features
- **Comprehensive Testing**: Unit, integration, and contract tests with >90% coverage
- **Clean Architecture**: Modular design with dependency injection and interfaces
- **Error Handling**: Robust error handling and graceful degradation
- **Process Safety**: Safe process management with proper cleanup and signal handling
- **Configurable**: Flexible configuration for different server types and environments

## Quick Start

### Prerequisites
- Go 1.24.1 or later
- Minecraft server executable
- Network access to target server (for RCON features)

### Installation
```bash
git clone <repository-url>
cd MinecraftServerManager
go mod tidy
```

### Running the Service
```bash
go run main.go
```

The service will start on port 8080 by default.

### Example Usage
```bash
# Start the server
curl -X POST http://localhost:8080/start

# Check server status
curl http://localhost:8080/status

# Stop the server
curl -X POST http://localhost:8080/stop
```

## Development

### Testing
```bash
# Run all tests
make all-test

# Run specific test types
make unit-test
make integration-test
make contract-test

# Generate coverage reports
make coverage-report
```

### Code Quality
```bash
# Run linter
make lint

# Build project
make build
```

### Project Structure
```
src/
├── controls/           # Core server control logic
├── httprouter/         # HTTP request routing and handling
├── httprouteradapter/  # HTTP adapter layer
├── os_api_adapter/     # OS-specific process operations
├── rcon/              # RCON client for server communication
├── remoteconnection/   # Remote connection management
└── windowsconstants/   # Windows-specific constants

tests/
├── integration/        # End-to-end integration tests
├── unit/              # Unit tests for individual components
└── contracts/         # Contract tests for external interfaces
```

## Configuration

The service can be configured for different environments:
- **Development**: Uses stub RCON adapter for testing
- **Production**: Connects to actual Minecraft server via RCON
- **Process Target**: Configurable server executable (defaults to notepad.exe for testing)

## Architecture

Built with clean architecture principles:
- **Separation of Concerns**: Clear boundaries between HTTP, business logic, and OS layers
- **Dependency Injection**: Flexible component composition for testing and configuration
- **Interface-Based Design**: Easy mocking and testing with well-defined contracts
- **Error Boundaries**: Proper error handling at each architectural layer

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass: `make all-test`
5. Run linting: `make lint`
6. Submit a pull request
