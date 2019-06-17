build:
	go install -v ./...
	go test -i -v ./...
	go test -c -v ./inventory
	go test -c -v ./management
	go test -c -v ./server

run:
	go test -v ./...

run-race:
	go test -v -race ./...
