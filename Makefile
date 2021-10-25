.PHONY: all auth game tidy

auth:
	go run cmd/auth/main.go

game:
	go run cmd/game/main.go

tidy:
	go mod tidy
