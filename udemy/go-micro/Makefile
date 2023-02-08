FRONT_END_BINARY=front
BROKER_BINARY=broker

# up: starts all containers without forcing build
start-back:
	@echo "Starting Docker images..."
	docker compose up
	@echo "Docker images started!"

# stops docker-compose (if running), builds all projects and starts docker compose
fresh-start-back: build-broker
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build
	@echo "Docker images built and started!"

# stop docker compose backend services
stop-back:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

# builds the broker binary as a Linux executable
build-broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

# builds the front end binary
build-front:
	@echo "Building front end binary..."
	cd ./front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

# starts the front end
start-front: build-front
	@echo "Starting front end"
	cd ./front-end && ./${FRONT_END_BINARY} &

# stops the front end (this is weird to have, because it stops the front end started and streaming logs on another tab)
stop-front:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"