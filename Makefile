SHELL := /bin/bash

.PHONY: clean

clean:
	rm -rf ./bin/

# Generate gRPC cli/server code for Go
.PHONY: protos containers bin

protos:
ifndef TARGET
	TARGET=all
endif
	$(MAKE) -f build/protos/Makefile $(TARGET)

containers:
	$(MAKE) -f build/container/Makefile TARGET=$(TARGET)

bin:
	$(MAKE) -f build/bins/Makefile TARGET=$(TARGET) all