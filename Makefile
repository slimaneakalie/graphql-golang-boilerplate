generate:
	go run github.com/99designs/gqlgen generate

build:
	GO111MODULE=on go build -o ./graphql-api.out cmd/service/*.go

run: build
	./graphql-api.out