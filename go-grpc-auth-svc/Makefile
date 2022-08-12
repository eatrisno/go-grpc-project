proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/*.proto

server:
	go run cmd/main.go
	
tidy:
	go mod tidy

run: 
	./bin/app

build:
	make proto
	make tidy
	go build -o ./bin/app ./cmd

develop:
	make proto
	make tidy
	make server

serve:
	make clean
	air

clean:
	rm -rf ./bin