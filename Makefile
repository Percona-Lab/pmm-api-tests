all:
	go install -v ./...
	go test -i -v ./...

run:
	go test -v ./...

run-race:
	go test -v -race ./...
