.PHONY: docker-build docker-dev clean

DOCKER_IMAGE_NAME := nausea-web
DOCKER_CONTAINER_NAME := nausea-web

docker-dev:
	@echo "Building and starting Docker containers..."
	@docker-compose up

docker-dev-build:
	@echo "Building and starting Docker containers..."
	@docker-compose up --build

docker-build:
	@echo "Building Docker image from Dockerfile..."
	@docker build -t $(DOCKER_IMAGE_NAME) .
	@echo "Running Docker container from image..."
	@docker run -d --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME)

clean:
	@echo "Stopping and removing Docker containers..."
	@docker-compose down
	@docker stop $(DOCKER_CONTAINER_NAME)
	@docker rm $(DOCKER_CONTAINER_NAME)
	@echo "Cleaning up Docker images and volumes..."
	@docker system prune -f
	@docker volume prune -f

check-dependencies:
	@command -v docker >/dev/null 2>&1 || { echo >&2 "Docker is not installed. Aborting."; exit 1; }
	@command -v docker-compose >/dev/null 2>&1 || { echo >&2 "Docker Compose is not installed. Aborting."; exit 1; }

all: check-dependencies docker-dev
