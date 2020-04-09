# Generate gRPC cli/server code for Go
.PHONY: all activity order report

all: activity order report

activity:
	protoc \
		-I protos/services \
		--go_out=plugins=grpc:pkg/services/activity/comm \
		protos/services/activity.proto

order:
	protoc \
		-I protos/services \
		--go_out=plugins=grpc:pkg/services/order/comm \
		protos/services/order.proto

report:
	protoc \
		-I protos/services \
		--go_out=plugins=grpc:pkg/services/report/comm \
		protos/services/report.proto
