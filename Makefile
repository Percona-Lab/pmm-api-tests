all:
	go install -v ./...
	go test -c -v ./inventory
	go test -c -v ./server
