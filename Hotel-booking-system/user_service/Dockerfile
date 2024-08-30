FROM golang:1.23.0-alpine AS builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o user_service cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/user_service .

EXPOSE 8080

CMD ["./user_service"]
