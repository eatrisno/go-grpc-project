proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/*.proto

server:
	go run main.go
	
tidy:
	go mod tidy

run: 
	./bin/order-svc

build:
	make proto
	make tidy
	go build -o ./bin/order-svc

develop:
	make proto
	make tidy
	make server

serve:
	make clean
	air

clean:
	rm -rf ./bin