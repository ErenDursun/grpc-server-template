BINARY_NAME=grpc-server
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
PB_VER="3.20.3"
export PATH:=${PATH}:${HOME}/.local/bin

all: protoc build

build:
	go mod tidy
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows cmd/${BINARY_NAME}/main.go

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/grpc/*/*/*.proto

tools:
	curl -LO ${PB_REL}/download/v${PB_VER}/protoc-${PB_VER}-linux-x86_64.zip
	unzip protoc-${PB_VER}-linux-x86_64.zip -d ${HOME}/.local
	rm protoc-${PB_VER}-linux-x86_64.zip
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
