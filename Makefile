BINARY_NAME=grpc-server
PB_VER="24.4"
PB_ZIP=protoc-${PB_VER}-linux-x86_64.zip
PB_REL="https://github.com/protocolbuffers/protobuf/releases/download/v${PB_VER}/${PB_ZIP}"
export PATH:=${PATH}:${HOME}/.local/bin

all: generate build

build:
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows.exe cmd/${BINARY_NAME}/main.go

generate:
	go generate ./...

tools:
	curl -LO ${PB_REL}
	unzip -o ${PB_ZIP} -d ${HOME}/.local bin/protoc
	unzip -o ${PB_ZIP} -d ${HOME}/.local 'include/*'
	rm ${PB_ZIP}
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
