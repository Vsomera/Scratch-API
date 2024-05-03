.DEFAULT_GOAL := help

build:
	@go build -o bin/main

run: build
	@./bin/main

help:
	@echo "Available commands:"
	@echo "  make run       : Build and run the API"
	@echo "  make build     : Build the API"
	@echo "  make show-help : Show this help message"