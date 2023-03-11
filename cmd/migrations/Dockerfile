FROM wokacz/hermod

RUN mkdir $GOPATH/src/github.com/wokacz/hermod/cmd/migrations

WORKDIR $GOPATH/src/github.com/wokacz/hermod

COPY . cmd/migrations/

RUN go build -o /bin/app-release ./cmd/migrations/main.go

ENTRYPOINT ["/bin/app-release"]