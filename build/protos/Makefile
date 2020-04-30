# Generate gRPC cli/server code for Go
.PHONY: all

fpath := protos/$(PRJ_TARGET)comm/$(PRJ_TARGET).proto

all: $(fpath)
	mkdir -p pkg/services/$(PRJ_TARGET)comm
	protoc \
		-I protos \
		--go_out=plugins=grpc:pkg/services \
		--go_opt=paths=source_relative \
		$(PRJ_TARGET)comm/$(PRJ_TARGET).proto

	mkdir -p front-end/front-office/gen/comm
	protoc \
		-I protos \
		--js_out=import_style=typescript,binary:front-end/front-office/src/gen/comm/$(PRJ_TARGET)comm \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:front-end/front-office/src/gen/comm \
		$(PRJ_TARGET)comm/$(PRJ_TARGET).proto