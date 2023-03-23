FROM wokacz/hermod

RUN mkdir $GOPATH/src/github.com/wokacz/hermod/cmd/roles

WORKDIR $GOPATH/src/github.com/wokacz/hermod

COPY . cmd/roles/

RUN go build -o /bin/app-release ./cmd/roles/main.go

ENTRYPOINT ["/bin/app-release"]