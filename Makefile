integration-test:
	go test .\tests\integration\godogs\...

unit-test:
	go test -v $(if $(dir),./tests/unit/$(dir),./tests/unit/...)

test-all: unit-test integration-test

build:
	go mod tidy