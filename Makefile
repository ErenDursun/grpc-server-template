BINARY_NAME=grpc-server

all: generate build

build:
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows.exe cmd/${BINARY_NAME}/main.go

generate:
	go generate ./...

tools:
	curl -fsS https://pkgx.sh | sh
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
