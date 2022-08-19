proto:
	protoc --go_out=. --go-grpc_out=. pkg/*/pb/*.proto

server:
	go run cmd/main.go
	
tidy:
	go mod tidy

run: 
	./bin/api-gateway

build:
	make proto
	make tidy
	go build -o ./bin/api-gateway ./cmd

develop:
	make proto
	make tidy
	make server

serve:
	make clean
	air

clean:
	rm -rf ./bin