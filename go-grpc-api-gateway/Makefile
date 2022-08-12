proto:
	protoc --go_out=. --go-grpc_out=. pkg/**/pb/*.proto

server:
	go run cmd/main.go

build:
	go build -o bin/app ./cmd

develop:
	make proto
	make server