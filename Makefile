.PHONY: all run build

all: run

run: build
	@echo "Running the application..."
	@./bin/main

build:
	@echo "Building the application..."
	@go build -o bin/main

test:
	@echo "Running tests..."
	@go test -v ./...

install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

docs:
	echo "Generating Swagger documentation..."
	@GO111MODULE=off swagger generate spec -o ./static/swagger.yaml --scan-models
