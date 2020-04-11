.PHONY: all cli srv

GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

ifndef PRJ_TARGET
$(error PRJ_TARGET is not set)
endif


all: cli srv

cli:
	mkdir -p bin/$(PRJ_TARGET)/
	$(GOMOD) verify
	$(GOGET) -v ./...

	CGO_ENABLED=0 $(GOBUILD) \
		-ldflags="-w -s" \
		-a \
		-installsuffix cgo \
		-o bin/$(PRJ_TARGET)/$(PRJ_TARGET)_cli \
		cmd/services/$(PRJ_TARGET)/client/main.go

srv:
	mkdir -p bin/$(PRJ_TARGET)/
	$(GOMOD) verify
	$(GOGET) -v ./...
	
	CGO_ENABLED=0 $(GOBUILD) \
		-ldflags="-w -s" \
		-a \
		-installsuffix cgo \
		-o bin/$(PRJ_TARGET)/$(PRJ_TARGET)_srv \
		cmd/services/$(PRJ_TARGET)/server/main.go