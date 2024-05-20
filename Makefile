APP_NAME = emporion

build:
	@go build -o bin/$(APP_NAME) cmd/main.go

run: build
	@./bin/$(APP_NAME)

test:
	@go test -v ./...