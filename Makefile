all: test vet build gen_docs

test:
	go test ./... -v

vet:
	go vet ./...

build:
	go build -o ./bin/rest_server ./cmd/rest

run:
	go run ./cmd/rest

gen_docs:
	swag init -g ./cmd/rest/rest.go
