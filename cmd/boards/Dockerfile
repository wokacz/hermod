FROM wokacz/hermod

RUN mkdir $GOPATH/src/github.com/wokacz/hermod/cmd/boards

WORKDIR $GOPATH/src/github.com/wokacz/hermod

COPY . cmd/boards/

RUN go build -o /bin/app-release ./cmd/boards/main.go

ENTRYPOINT ["/bin/app-release"]