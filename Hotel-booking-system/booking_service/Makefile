run:
	go run cmd/main.go
table:
	migrate create -dir migrations -ext sql db
up:
	migrate -path migrations -database "postgres://postgres:14022014@localhost/booking_service?sslmode=disable" up
down:
	migrate -path migrations -database "postgres://postgres:14022014@localhost/booking_service?sslmode=disable" down
force:
	migrate -path migrations -database "postgres://postgres:14022014@localhost/booking_service?sslmode=disable" force 