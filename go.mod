module github.com/Percona-Lab/pmm-api-tests

go 1.14

// Use for local development, but do not commit:
// replace github.com/percona/pmm => ../../pmm

// Update with:
// go get -v github.com/percona/pmm@PMM-2.0

require (
	github.com/AlekSi/pointer v1.1.0
	github.com/brianvoe/gofakeit/v5 v5.11.0
	github.com/davecgh/go-spew v1.1.1
	github.com/go-openapi/runtime v0.19.20
	github.com/percona/pmm v2.12.1-0.20201202114955-5089534cb973+incompatible
	github.com/prometheus/client_golang v1.7.1
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20200722175500-76b94024e4b6
	google.golang.org/grpc v1.30.0
)
