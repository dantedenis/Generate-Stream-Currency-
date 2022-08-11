BINARY_NAME=generate_service

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ./cmd/main.go

run:
	./${BINARY_NAME}

all: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin