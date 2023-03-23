FROM wokacz/hermod

RUN mkdir $GOPATH/src/github.com/wokacz/hermod/cmd/notifications

WORKDIR $GOPATH/src/github.com/wokacz/hermod

COPY . cmd/notifications/

RUN go build -o /bin/app-release ./cmd/notifications/main.go

ENTRYPOINT ["/bin/app-release"]