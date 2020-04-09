SHELL := /bin/bash

.PHONY: all clean

clean:
	rm -rf ./bin/

##################################################
#################### SERVICES ####################
##################################################

# Generate gRPC cli/server code for Go
.PHONY: protos

protos:
ifndef TARGET
	TARGET=all
endif
	$(MAKE) -f build/protos/Makefile $(TARGET)
