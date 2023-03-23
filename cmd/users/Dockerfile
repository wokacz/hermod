FROM wokacz/hermod

RUN mkdir $GOPATH/src/github.com/wokacz/hermod/cmd/users

WORKDIR $GOPATH/src/github.com/wokacz/hermod

COPY . cmd/users/

RUN go build -o /bin/app-release ./cmd/users/main.go

ENTRYPOINT ["/bin/app-release"]