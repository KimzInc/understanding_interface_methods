.DEFAULT_GOAL := build

.PHONY: fmt vet build run clean

fmt:
	@echo "Formatting Go files..."
	@go fmt ./...

vet: fmt
	@echo "Vetting Go files..."
	@go vet ./...

build: vet
	@echo "Building Go application..."
	@go build -o main .

run: build
	@echo "Running Go application..."
	@./main

clean:
	@echo "Cleaning up..."
	@rm -f main