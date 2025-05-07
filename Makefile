integration-test:
	go test ./tests/integration/godogs/... -v

unit-test:
	go test $(if $(dir),./tests/unit/$(dir),./tests/unit/...) -v

all-test: unit-test integration-test

build:
	go mod tidy