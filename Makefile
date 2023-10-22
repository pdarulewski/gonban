BINARY_NAME=gonban
APP_PATH=./cmd/gonban


all: build

build:
	go build -o $(BINARY_NAME) $(APP_PATH)

run: build
	go run $(APP_PATH)
