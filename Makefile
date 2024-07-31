
all: dirs pitch length import_message boundary note rest symbol \
		accidental embellishment movement tie \
		timeline tuplet barline time_signature \
		measure tune

clean:
	find . -name *.pb.go -delete

build:
	buf generate

lint:
	buf lint
