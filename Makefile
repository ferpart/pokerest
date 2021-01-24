default: help

test: ## tests all files in directory
	go test ./...

build: ## builds go file
	go build

run: ## runs without building
	go run main.go