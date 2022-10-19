ifeq ($(OS),Windows_NT)
    REMOVE_DIRECTORY_CMD = rmdir /q /s
    REMOVE_CMD = del
    BINARY_NAME = main.exe
    BINARY_DEST = c:/windows/system32/config_master
else
    REMOVE_DIRECTORY_CMD = rm -rf
    REMOVE_CMD = rm
    BINARY_NAME = main
    BINARY_DEST = /usr/local/bin/config_master
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

install:
	cp bin/server/${BINARY_NAME} ${BINARY_DEST}

uninstall:
	${REMOVE_CMD} ${BINARY_DEST}