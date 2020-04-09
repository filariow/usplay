# build executables
.PHONY: all cli svr

ifndef TARGET
$(error TARGET is not set)
endif

all: cli svr

cli:
	mkdir -p bin/$(TARGET)/
	go mod verify
	go get -v ./...
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo -o bin/$(TARGET)/cli_$(TARGET) cmd/services/$(TARGET)/client/main.go

svr:
	mkdir -p bin/activity/
	go mod verify
	go get -v ./...
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo -o bin/$(TARGET)/svr_$(TARGET) cmd/services/$(TARGET)/server/main.go