ifeq ($(OS),Windows_NT)
    REMOVE_DIRECTORY_CMD = rmdir /q /s
else
    REMOVE_DIRECTORY_CMD = rm -rf
endif

all: test vet build run

test:
	go test -cover ./...

vet:
	go vet ./...

build:
	go build -o bin/server/main.exe cmd/server/main.go

clean:
	${REMOVE_DIRECTORY_CMD} bin

run: build
	bin/server/main.exe