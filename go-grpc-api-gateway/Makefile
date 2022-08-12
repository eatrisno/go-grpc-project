proto:
	protoc --go_out=. --go-grpc_out=. pkg/*/pb/*.proto

server:
	go run cmd/main.go
	
tidy:
	go mod tidy

run: 
	./bin/app

build:
	make tidy
	make proto
	go build -o ./bin/app ./cmd

develop:
	make tidy
	make proto
	make server

serve:
	make clean
	air

clean:
	rm -rf ./bin