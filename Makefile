# BACKEND_BINARY=backendApp

up:
	@echo "Starting containers..."
	docker compose up -d
	@echo "Containers started!"

up_build:
	@echo "Stopping containers"
	docker compose down
	@echo "Building and starting containers..."
	docker compose up --build -d
	@echo "Docker images built and containers started!"

down:
	@echo "Stopping docker containers"
	docker compose down
	@echo "Done!"


## build_backend:
## 	@echo "Building backend binary..."
## 	cd ./backend && env GOOS=linux CGO_ENABLED=0 go build -o ${BACKEND_BINARY} ./cmd/api
## 	@echo "Done!"
