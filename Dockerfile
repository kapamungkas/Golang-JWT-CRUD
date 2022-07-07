FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go build -o betest ./cmd/cmd_serve.go

ENTRYPOINT ["/app/betest"]