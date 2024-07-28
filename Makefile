
OUT_DIR=musicmodel

all: dirs pitch length note rest symbol

dirs:
	mkdir -p "$(OUT_DIR)" && \
	mkdir -p "$(OUT_DIR)/pitch" && \
	mkdir -p "$(OUT_DIR)/length" && \
	mkdir -p "$(OUT_DIR)/symbols"

clean:
	find . -name *.pb.go -delete

pitch:
	protoc --go_out=$(OUT_DIR)/pitch --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR)/pitch --go-grpc_opt=paths=source_relative \
		pitch.proto
length:
	protoc --go_out=$(OUT_DIR)/length --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR)/length --go-grpc_opt=paths=source_relative \
		length.proto

note:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbols/note.proto

rest:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbols/rest.proto

symbol:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbol.proto