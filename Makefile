all: build

init:           ## Installs development tools
	go build -modfile=tools/go.mod -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
	go build -modfile=tools/go.mod -o bin/go-junit-report github.com/jstemmer/go-junit-report
	go build -modfile=tools/go.mod -o bin/reviewdog github.com/reviewdog/reviewdog/cmd/reviewdog


build:
	go install -v ./...
	go test -v ./...
	go test -c -v ./inventory
	go test -c -v ./management
	go test -c -v ./server

dev-test:						## Run test on dev env. Use `PMM_KUBECONFIG=/path/to/kubeconfig.yaml make dev-test` to run tests for DBaaS.
	go test -count=1 -p 1 -v ./... -pmm.server-insecure-tls

run:
	go test -count=1 -p 1 -v ./... 2>&1 | tee pmm-api-tests-output.txt
	cat pmm-api-tests-output.txt | bin/go-junit-report > pmm-api-tests-junit-report.xml

run-race:
	go test -count=1 -p 1 -v -race ./... 2>&1 | tee pmm-api-tests-output.txt
	cat pmm-api-tests-output.txt | bin/go-junit-report > pmm-api-tests-junit-report.xml

FILES = $(shell find . -type f -name '*.go')

format:                         ## Format source code.
	gofmt -w -s $(FILES)
	goimports -local github.com/Percona-Lab/pmm-api-tests -l -w $(FILES)

clean:
	rm -f ./pmm-api-tests-output.txt
	rm -f ./pmm-api-tests-junit-report.xml

check-all:                      ## Run golang ci linter to check new changes from master.
	bin/golangci-lint run -c=.golangci.yml --new-from-rev=master

ci-reviewdog:                   ## Runs reviewdog checks.
	bin/golangci-lint run -c=.golangci-required.yml --out-format=line-number | bin/reviewdog -f=golangci-lint -level=error -reporter=github-pr-check
	bin/golangci-lint run -c=.golangci.yml --out-format=line-number | bin/reviewdog -f=golangci-lint -level=error -reporter=github-pr-review
