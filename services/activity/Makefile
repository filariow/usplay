.PHONY: proto

proto:
	mkdir -p gen

	protoc \
		-I protos/ \
		--go_out=plugins=grpc:gen \
		protos/*.proto
	