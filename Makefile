integration-test:
	go test ./tests/integration/godogs/... -v

contract-test:
	go test ./tests/integration/contracts/... -v

unit-test:
	go test $(if $(dir),./tests/unit/$(dir),./tests/unit/...) -v

all-test: unit-test integration-test contract-test

build:
	go mod tidy

# Coverage directory
coverage-dir:
	if not exist coverage mkdir coverage

# Package-specific coverage targets
cover-controls: coverage-dir
	go test .\tests\... -coverpkg=minecraftremote/src/controls -covermode=atomic -v -coverprofile=coverage\controls.out

cover-process: coverage-dir
	go test .\tests\... -coverpkg=minecraftremote/src/process -covermode=atomic -v -coverprofile=coverage\process.out

cover-rcon: coverage-dir
	go test .\tests\... -coverpkg=minecraftremote/src/rcon -covermode=atomic -v -coverprofile=coverage\rcon.out

cover-httprouter: coverage-dir
	go test .\tests\... -coverpkg=minecraftremote/src/httprouter -covermode=atomic -v -coverprofile=coverage\httprouter.out

cover-httprouteradapter: coverage-dir
	go test .\tests\... -coverpkg=minecraftremote/src/httprouteradapter -covermode=atomic -v -coverprofile=coverage\httprouteradapter.out

cover-remoteconnection: coverage-dir
	go test .\tests\... -coverpkg=minecraftremote/src/remoteconnection -covermode=atomic -v -coverprofile=coverage\remoteconnection.out

# Generate HTML coverage reports
coverage-html: coverage-dir
	go tool cover -html=coverage\controls.out -o coverage\controls.html
	go tool cover -html=coverage\process.out -o coverage\process.html
	go tool cover -html=coverage\rcon.out -o coverage\rcon.html
	go tool cover -html=coverage\httprouter.out -o coverage\httprouter.html
	go tool cover -html=coverage\httprouteradapter.out -o coverage\httprouteradapter.html
	go tool cover -html=coverage\remoteconnection.out -o coverage\remoteconnection.html

# Run all coverage tests
cover-all: cover-controls cover-process cover-rcon cover-httprouter cover-httprouteradapter cover-remoteconnection coverage-html

# Clean coverage files
clean-coverage:
	powershell -ExecutionPolicy Bypass -Command "if (Test-Path coverage) { Remove-Item -Path coverage -Recurse -Force }"

# Merge coverage profiles
merge-coverage: coverage-dir
	powershell -ExecutionPolicy Bypass -Command "Get-Content coverage/*.out | Select-String -NotMatch 'mode: atomic' | Set-Content coverage/merged.txt"
	powershell -ExecutionPolicy Bypass -Command "Get-Content coverage/controls.out | Select-String 'mode: atomic' -SimpleMatch | Set-Content -Path coverage/merged.out"
	powershell -ExecutionPolicy Bypass -Command "Get-Content coverage/merged.txt | Add-Content -Path coverage/merged.out"
	powershell -ExecutionPolicy Bypass -Command "Remove-Item -Path coverage/merged.txt -Force"

# Generate single HTML report from merged profiles
coverage-report: cover-all merge-coverage
	go tool cover -html=coverage/merged.out -o coverage/full_coverage.html


lint:
	golangci-lint run ./...

.PHONY: coverage-dir cover-controls cover-process cover-rcon cover-httprouter cover-httprouteradapter cover-remoteconnection coverage-html cover-all clean-coverage merge-coverage coverage-report lint