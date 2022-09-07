proto:
	protoc --go_out=. --go-grpc_out=. pkg/*/pb/*.proto

server:
	go run main.go
	
tidy:
	go mod tidy

run: 
	./bin/api-gateway

build:
	swag init
	make proto
	make tidy
	go build -o ./bin/api-gateway

develop:
	swag init
	make proto
	make tidy
	make server

serve:
	make clean
	air

clean:
	rm -rf ./bin