start:
	docker compose up --build -d

start_%:
	docker compose up -d --build $*

start_debug:
	docker compose up --build

restart: stop start

restart_%:
	docker compose stop $*
	docker compose rm -f $*
	docker compose up --build -d $*

stop:
	docker compose down --remove-orphans

stop_%:
	docker compose stop $*
	docker compose rm -f $*

logs:
	docker compose logs -f

build: build_base build_migrations build_authorization

build_base:
	@echo "\033[92mBuilding image wokacz/hermod\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod -f ./Dockerfile ./

build_migrations:
	@echo "\033[92mBuilding image wokacz/hermod-migrations\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod-migrations -f ./cmd/migrations/Dockerfile ./cmd/migrations/

build_authorization:
	@echo "\033[92mBuilding image wokacz/hermod-authorization\033[0m"
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod-authorization -f ./cmd/authorization/Dockerfile ./cmd/authorization/
