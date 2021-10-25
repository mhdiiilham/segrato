.PHONY: all

run-auth:
	go run cmd/auth/main.go

tidy:
	go mod tidy
