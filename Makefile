ifeq ($(OS),Windows_NT)
    REMOVE_DIRECTORY_CMD = rmdir /q /s
    BINARY_NAME = main.exe
else
    REMOVE_DIRECTORY_CMD = rm -rf
    BINARY_NAME = main
endif

all: clean validate run

test:
	go test -cover ./...

vet:
	go vet ./...

validate: test vet

build:
	go build -o bin/server/${BINARY_NAME} cmd/server/main.go

clean:
	${REMOVE_DIRECTORY_CMD} bin

run: build
	bin/server/${BINARY_NAME}