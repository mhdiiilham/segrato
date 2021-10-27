.PHONY: all auth game tidy test

auth:
	go run cmd/auth/main.go

game:
	go run cmd/game/main.go

tidy:
	go mod tidy

test:
	go test -v -race -cover ./...
