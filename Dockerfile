FROM golang:1.18-alpine AS builder

WORKDIR /src
COPY go.mod .
COPY go.sum .
COPY *.go ./
COPY server luckless_server
RUN go build -o query luckless_server/main.go

FROM alpine

WORKDIR /app
COPY --from=builder /src/query .

CMD ./query