export GO111MODULE=on

all: deps build
build:
	go build -o tlfgen cmd/tlfgen/main.go
	go build -o tlfserv cmd/tlfserv/main.go
clean:
	go clean
	rm -f tlfgen tlfserv
deps:
	go build -v ./...
upgrade:
	go get -u
