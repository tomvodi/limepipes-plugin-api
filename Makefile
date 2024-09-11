
clean:
	find . -name *.pb.go -delete

build:
	buf generate

lint:
	buf lint && golangci-lint run

mocks:
	mockery
