.PHONY: all auth game tidy test mock-prepare mock

build:
	go build -o auth-service cmd/auth/main.go
	go build -o game-service cmd/game/main.go

auth:
	go run cmd/auth/main.go

game:
	go run cmd/game/main.go

tidy:
	go mod tidy

test:
	go test -race -cover ./...

mock-prepare:
	go install github.com/golang/mock/mockgen@v1.6.0
	go get -u github.com/golang/mock/gomock
	go get -u github.com/bxcodec/faker/v3

mock:
	mockgen -source=pkg/password/interface.go -destination=pkg/password/mock/interface_mock.go -package=mock
	mockgen -source=pkg/token/interface.go -destination=pkg/token/mock/token_interface_mock.go -package=mock
	mockgen -source=user/interface.go -destination=user/mock/user_interface_mock.go -package=mock
