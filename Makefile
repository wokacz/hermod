# golang

go_mod:
	go mod download -x
	go mod vendor -v

# docker
build: build_base build_migrations build_authorization

build_base:
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod -f ./Dockerfile ./

build_migrations:
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod/migrations -f ./cmd/migrations/Dockerfile ./cmd/migrations/

build_authorization:
	DOCKER_BUILDKIT=1 docker build -t wokacz/hermod/authorization -f ./cmd/authorization/Dockerfile ./cmd/authorization/




