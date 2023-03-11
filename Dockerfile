FROM golang:1.20-alpine AS builder

# enable Go modules support
#ENV GO111MODULE=on

# All these steps will be cached
WORKDIR $GOPATH/src/github.com/wokacz/hermod

# COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .

RUN mkdir cmd