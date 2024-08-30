FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hotel-service ./cmd/main.go

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/booking_service .

EXPOSE 8082

CMD ["/home/shakhboz/go/src/5.project/booking_service"]
