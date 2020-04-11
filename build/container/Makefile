# Build containers for client and server	
.PHONY: all cli srv cli.raw cli.ubuntu srv.raw srv.ubuntu

ifndef PRJ_TARGET
$(error PRJ_TARGET is not defined)
endif

ifndef TARGET
$(error TARGET is not defined)
endif

all: srv cli
cli: cli.raw cli.ubuntu
srv: srv.raw srv.ubuntu

cli.raw:
	docker build \
		--build-arg USVC_NAME=$(PRJ_TARGET) \
		-t usplay/$(PRJ_TARGET)_cli:"$$(date +%s)" \
		-t usplay/$(PRJ_TARGET)_cli:raw-"$$(date +%s)" \
		-f build/container/usvc/client/Dockerfile \
		.

cli.ubuntu:
	docker build \
		--build-arg USVC_NAME=$(PRJ_TARGET) \
		-t usplay/$(PRJ_TARGET)_cli:ub-"$$(date +%s)" \
		-f build/container/usvc/client/ubuntu.Dockerfile \
		.

srv.raw:
	docker build \
		--build-arg USVC_NAME=$(PRJ_TARGET) \
		-t usplay/$(PRJ_TARGET)_srv:"$$(date +%s)" \
		-t usplay/$(PRJ_TARGET)_srv:raw-"$$(date +%s)" \
		-f build/container/usvc/server/Dockerfile \
		.

srv.ubuntu:
	docker build \
		--build-arg USVC_NAME=$(PRJ_TARGET) \
		-t usplay/$(PRJ_TARGET)_srv:ub-"$$(date +%s)" \
		-f build/container/usvc/server/ubuntu.Dockerfile \
		.
