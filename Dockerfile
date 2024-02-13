FROM golang:1.20.14-alpine

COPY go.mod go.sum /go/src/github.com/franzinBr/feedks-api/
WORKDIR /go/src/github.com/franzinBr/feedks-api
RUN go mod download
COPY . /go/src/github.com/franzinBr/feedks-api
RUN go build -o /usr/bin/feedks-api ./cmd/main.go

EXPOSE 3000 3000
ENTRYPOINT ["/usr/bin/feedks-api"]