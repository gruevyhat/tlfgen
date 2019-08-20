export GO111MODULE=on

all: deps build
install:
	go install cmd/tlfgen/main.go
	go install cmd/tlfserv/main.go
build:
	go build -o tlfgen cmd/tlfgen/main.go
	go build -o tlfserv cmd/tlfserv/main.go
test:
	go test
clean:
	go clean
	rm -f tlfgen tlfserv
deps:
	go build -v ./...
upgrade:
	go get -u
