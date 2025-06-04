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

build:
	docker-compose build --no-cache

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

restart: down build up

logs:
	docker-compose logs -f

ps:
	docker ps
