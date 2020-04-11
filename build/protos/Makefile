# Generate gRPC cli/server code for Go
.PHONY: all

fpath := protos/services/$(PRJ_TARGET).proto

all: $(fpath)
	mkdir -p pkg/services/$(PRJ_TARGET)/comm
	protoc \
		-I protos/services \
		--go_out=plugins=grpc:pkg/services/$(PRJ_TARGET)/comm \
		protos/services/$(PRJ_TARGET).proto
