proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/*.proto

server:
	go run main.go
	
tidy:
	go mod tidy

run: 
	./bin/auth-svc

build:
	make proto
	make tidy
	go build -o ./bin/auth-svc

init:
	mkdir -p ~/go/src/github.com/eatrisno
	ln -sf `pwd` ~/go/src/github.com/eatrisno/go-grpc-auth-svc
develop:
	make proto
	make tidy
	make server

serve:
	make clean
	air

clean:
	rm -rf ./bin