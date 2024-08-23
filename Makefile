
clean:
	find . -name *.pb.go -delete

build:
	buf generate

lint:
	buf lint

mocks:
	mockery
