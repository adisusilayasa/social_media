ping:
	echo "Pong"

run:
	go run ./cmd/app/main.go

swaggo:
	swag init -g ./cmd/app/main.go --output docs 