SHELL := /bin/bash

.PHONY: clean

clean:
	rm -rf ./bin/

.PHONY: protos container bin

PROJECTS := activity report order

protos:
ifndef PRJ_TARGET
	$(error PRJ_TARGET is not defined)
endif
	$(MAKE) -f build/protos/Makefile PRJ_TARGET=$(PRJ_TARGET)

container:
ifndef TARGET
	TARGET=all
endif
	$(MAKE) -f build/container/Makefile $(TARGET) PRJ_TARGET=$(PRJ_TARGET)

bin: 
ifndef TARGET
	TARGET=all
endif
	$(MAKE) -f build/bin/Makefile $(TARGET) PRJ_TARGET=$(PRJ_TARGET)
