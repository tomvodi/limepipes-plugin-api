
OUT_DIR=musicmodel

all: dirs pitch length boundary note rest symbol \
		accidental embellishment movement tie \
		timeline tuplet barline time_signature

dirs:
	mkdir -p "$(OUT_DIR)" && \
	mkdir -p "$(OUT_DIR)/pitch" && \
	mkdir -p "$(OUT_DIR)/length" && \
	mkdir -p "$(OUT_DIR)/symbols" && \
	mkdir -p "$(OUT_DIR)/boundary" && \
	mkdir -p "$(OUT_DIR)/barline"


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

boundary:
	protoc --go_out=$(OUT_DIR)/boundary --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR)/boundary --go-grpc_opt=paths=source_relative \
		boundary.proto

time_signature:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		time_signature.proto

barline:
	protoc --go_out=$(OUT_DIR)/barline --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR)/barline --go-grpc_opt=paths=source_relative \
		barline.proto

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

accidental:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbols/accidental/accidental.proto

embellishment:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbols/embellishment/embellishment.proto

movement:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbols/movement/movement.proto

tie:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbols/tie/tie.proto

timeline:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbols/timeline/timeline.proto

tuplet:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		symbols/tuplet/tuplet.proto
