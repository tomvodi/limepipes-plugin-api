GOBIN ?= $$(go env GOPATH)/bin

clean:
	find . -name *.pb.go -delete

build:
	buf generate

lint:
	buf lint && golangci-lint run

mocks:
	mockery

.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yaml