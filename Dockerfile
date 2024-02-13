FROM golang:1.20.14-alpine

COPY go.mod go.sum /go/src/github.com/franzinBr/feedks-api/
WORKDIR /go/src/github.com/franzinBr/feedks-api
RUN go mod download
COPY . /go/src/github.com/franzinBr/feedks-api
RUN go build -o /usr/bin/feedks-api github.com/franzinBr/feedks-api/cmd/main.go

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/feedks-api"]