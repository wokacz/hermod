FROM wokacz/hermod

RUN mkdir $GOPATH/src/github.com/wokacz/hermod/cmd/app

WORKDIR $GOPATH/src/github.com/wokacz/hermod

COPY . cmd/app/

RUN go build -o /bin/app-release cmd/app/main.go

ENTRYPOINT ["/bin/app-release"]