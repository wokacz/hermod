FROM wokacz/hermod

RUN mkdir $GOPATH/src/github.com/wokacz/hermod/cmd/TO_REPLACE

WORKDIR $GOPATH/src/github.com/wokacz/hermod

COPY . cmd/TO_REPLACE/

RUN go build -o /bin/app-release ./cmd/TO_REPLACE/main.go

ENTRYPOINT ["/bin/app-release"]