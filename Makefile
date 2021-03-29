all:
	@echo "no default"

# tests
.PHONY: test
test:
	@go test -v tests/*.go

# build
.PHONY: build
build:
	@go build .
