.PHONY: all

fpath := protos/$(PRJ_TARGET)grpc/*.proto

all: $(fpath)
	mkdir -p pkg/services/$(PRJ_TARGET)grpc
	protoc \
		-I protos \
		--go_out=plugins=grpc:pkg/services \
		--go_opt=paths=source_relative \
		$(fpath)

	mkdir -p front-end/front-office/src/gen
	protoc \
		--plugin=protoc-gen-ts=./front-end/front-office/node_modules/.bin/protoc-gen-ts \
		--js_out=import_style=commonjs,binary:front-end/front-office/src/gen \
		--ts_out=service=grpc-web:front-end/front-office/src/gen \
		-I ./protos \
		$(fpath)
