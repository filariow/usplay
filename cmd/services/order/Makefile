# Generate gRPC cli/server code for Go
.PHONY: proto

proto:
	mkdir -p gen

	protoc \
		-I protos/ \
		--go_out=plugins=grpc:gen \
		protos/*.proto

# Build containers for client and server	
.PHONY: bcall bccli bccli.raw bccli.ubuntu bcsvr bcsvr.raw bcsvr.ubuntu

bcall: bcsvr bccli
bccli: bcli.raw bcli.ubuntu

bccli.raw:
	docker build -t usplay/us_order:"$$(date +%s)" -f build/container/client/Dockerfile .

bccli.ubuntu:
	docker build -t usplay/us_order:"$$(date +%s)" -f build/container/client/ubuntu.Dockerfile .

bcsvr: bsvr.raw bsvr.ubuntu

bcsvr.raw:
	docker build -t usplay/us_order:"$$(date +%s)" -f build/container/server/Dockerfile .

bcsvr.ubuntu:
	docker build -t usplay/us_order:"$$(date +%s)" -f build/container/server/ubuntu.Dockerfile .

# build executables
.PHONY: ball bcli bsvr

ball: bcli bsvr

bcli:
	mkdir -p bin
	go mod verify
	go get -v ./...
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo -o bin/cli_order cmd/client/main.go

bsvr:
	mkdir -p bin
	go mod verify
	go get -v ./...
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo -o bin/svr_order cmd/server/main.go