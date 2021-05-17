FROM golang:1.16

WORKDIR /go/src/github.com/luckless-finance/query
COPY go.mod .
COPY go.sum .
COPY *.go ./
COPY server luckless_server

CMD go run luckless_server/main.go