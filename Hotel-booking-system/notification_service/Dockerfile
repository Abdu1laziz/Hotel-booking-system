FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notification_service ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/notification_service /usr/local/bin/users-service

COPY certs /app/certs

EXPOSE 8083

CMD ["/home/shakhboz/go/src/5.project/notification_service"]
