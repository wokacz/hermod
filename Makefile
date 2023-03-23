# Start all containers
start:
	docker compose up --build -d

# Start container
start_%: build_base
	docker compose up -d --build $*

# Start all containers in debug mode
start_debug:
	docker compose up --build

# Restart all containers
restart: stop start

# Restart container
restart_%: build_base
	docker compose stop $*
	docker compose rm -f $*
	docker compose up --build -d $*

# Stop all containers
stop:
	docker compose down --remove-orphans

# Stop and remove container
stop_%:
	docker compose stop $*
	docker compose rm -f $*

# Show logs
logs:
	docker compose logs -f

# Build all images
build: build_base build_authorization build_roles build_users build_boards build_notifications

# Build base image
build_base:
	@echo "\033[92mBuilding image wokacz/hermod\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod -f ./Dockerfile ./

# Build authorization image
build_authorization:
	@echo "\033[92mBuilding image wokacz/hermod-authorization\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod-authorization -f ./cmd/authorization/Dockerfile ./cmd/authorization/

# Build users image
build_roles:
	@echo "\033[92mBuilding image wokacz/hermod-users\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod-users -f ./cmd/users/Dockerfile ./cmd/users/

# Build users image
build_users:
	@echo "\033[92mBuilding image wokacz/hermod-users\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod-users -f ./cmd/users/Dockerfile ./cmd/users/

# Build boards image
build_boards:
	@echo "\033[92mBuilding image wokacz/hermod-boards\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod-boards -f ./cmd/boards/Dockerfile ./cmd/boards/

build_notifications:
	@echo "\033[92mBuilding image wokacz/hermod-notifications\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod-notifications -f ./cmd/notifications/Dockerfile ./cmd/notifications/