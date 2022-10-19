BIN_DIR = bin

ifeq ($(OS),Windows_NT)
    REMOVE_DIRECTORY_CMD = if exist ${BIN_DIR} rmdir /q /s ${BIN_DIR}
    REMOVE_CMD = del
    BINARY_NAME = main.exe
    BINARY_DEST = c:/windows/system32/config_master
else
    REMOVE_DIRECTORY_CMD = rm -rf ${BIN_DIR} || true
    REMOVE_CMD = rm
    BINARY_NAME = main
    BINARY_DEST = /usr/local/bin/config_master
endif

all: clean validate build

test:
	go test -cover ./...

vet:
	go vet ./...

validate: test vet

build:
	go build -o bin/server/${BINARY_NAME} cmd/server/main.go

clean:
	${REMOVE_DIRECTORY_CMD}

run: build
	bin/server/${BINARY_NAME}

install:
	cp bin/server/${BINARY_NAME} ${BINARY_DEST}

uninstall:
	${REMOVE_CMD} ${BINARY_DEST}