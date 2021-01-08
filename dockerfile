FROM golang:latest
WORKDIR $GOPATH/src/github.com/demo/api
COPY . $GOPATH/src/github.com/demo/api
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./mango-api"]