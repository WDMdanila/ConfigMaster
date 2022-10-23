BIN_DIR = bin

ifeq ($(OS),Windows_NT)
    REMOVE_DIRECTORY_CMD = if exist ${BIN_DIR} rmdir /q /s ${BIN_DIR}
    BINARY_NAME = main.exe
    BINARY_DEST = c:/windows/system32/config_master
    REMOVE_CMD = if exist ${BINARY_DEST} del ${BINARY_DEST}
else
    REMOVE_DIRECTORY_CMD = rm -rf ${BIN_DIR} || true
    BINARY_NAME = main
    BINARY_DEST = /usr/local/bin/config_master
    REMOVE_CMD = rm ${BINARY_DEST} || true
endif

.PHONY: all
all: clean validate build

test:
	go test -cover ./...

vet:
	go vet ./...

validate: test vet

build:
	go build -o bin/server/${BINARY_NAME} cmd/server/main.go

.PHONY: clean
clean:
	${REMOVE_DIRECTORY_CMD}

run: build
	bin/server/${BINARY_NAME}

.PHONY: install
install:
	cp bin/server/${BINARY_NAME} ${BINARY_DEST}

uninstall:
	${REMOVE_CMD}

.PHONY: docker
docker:
	docker build -t config_master -f docker/Dockerfile .