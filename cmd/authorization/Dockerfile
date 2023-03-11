FROM wokacz/hermod

RUN mkdir $GOPATH/src/github.com/wokacz/hermod/cmd/authorization

WORKDIR $GOPATH/src/github.com/wokacz/hermod

COPY . cmd/authorization/

RUN go build -o /bin/app-release cmd/authorization/main.go

ENTRYPOINT ["/bin/app-release"]