# Build containers for client and server	
.PHONY: all cli cli.raw cli.ubuntu svr svr.raw svr.ubuntu

ifndef TARGET
$(error TARGET is not defined)
endif

all: svr cli
cli: cli.raw cli.ubuntu
svr: svr.raw svr.ubuntu

cli.raw:
	docker build -t usplay/us_$(TARGET):"$$(date +%s)" -f build/container/$(TARGET)/client/Dockerfile .

cli.ubuntu:
	docker build -t usplay/us_$(TARGET):"$$(date +%s)" -f build/container/$(TARGET)/client/ubuntu.Dockerfile .

svr.raw:
	docker build -t usplay/us_$(TARGET):"$$(date +%s)" -f build/container/$(TARGET)/server/Dockerfile .

svr.ubuntu:
	docker build -t usplay/us_$(TARGET):"$$(date +%s)" -f build/container/$(TARGET)/server/ubuntu.Dockerfile .
