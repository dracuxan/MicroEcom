.PHONY: all run build

all: run

run: build
	@echo "Running the application..."
	@./bin/main

test:
	@echo "Running tests..."
	@go test -v ./...

install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

docs: 
	echo "Generating Swagger documentation..."
	@GO111MODULE=off swagger generate spec -o ./static/swagger.yaml --scan-models

BINARY_NAME=microecom
BUILD_DIR=build
DEPLOY_DIR=/var/www/microecom

build: docs
	@echo "Building the Go binary..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

deploy: build
	@echo "Creating deployment directory..."
	sudo mkdir -p $(DEPLOY_DIR)
	sudo cp -r static/swagger.yaml $(DEPLOY_DIR)/static

	@echo "Stopping running service..."
	sudo systemctl stop microecom

	@echo "Deploying binary..."
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

	@echo "Restarting service..."
	sudo systemctl start microecom

clean:
	rm -rf $(BUILD_DIR)
