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

	mkdir -p front-end/front-office/src/gen
	# protoc \
	# 	-I protos \
	# 	--js_out=import_style=commonjs,binary:front-end/front-office/src/gen/comm/ \
	# 	--grpc-web_out=import_style=typescript,mode=grpcwebtext:front-end/front-office/src/gen/comm \
	# 	$(PRJ_TARGET)comm/$(PRJ_TARGET).proto

	protoc \
		--plugin=protoc-gen-ts=./front-end/front-office/node_modules/.bin/protoc-gen-ts \
		--js_out=import_style=commonjs,binary:front-end/front-office/src/gen \
		--ts_out=service=grpc-web:front-end/front-office/src/gen \
		-I ./protos \
		$(PRJ_TARGET)comm/$(PRJ_TARGET).proto

	# protoc -I=./protos $(PRJ_TARGET)comm/$(PRJ_TARGET).proto \
	# 	--js_out=import_style=commonjs,binary:front-end/front-office/src/gen \
	# 	--grpc-web_out=import_style=typescript,mode=grpcwebtext:front-end/front-office/src/gen

	# protoc --proto_path=protos \
	# 	--js_out=import_style=commonjs,binary:front-end/front-office/src/gen \
	# 	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:front-end/front-office/src/gen \
	# 	--ts_out=front-end/front-office/src/gen \
	# 	$(PRJ_TARGET)comm/$(PRJ_TARGET).proto

