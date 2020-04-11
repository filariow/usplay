############################
# STEP 1 build executable binary
############################
FROM golang:1.14.1-buster as builder
ARG USVC_NAME

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR $GOPATH/src/github.com/FrancescoIlario/usplay

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download \
    && go mod verify

COPY . .

# Build the binary
RUN rm -rf ./bin \
    && make bin PRJ_TARGET=${USVC_NAME} TARGET=srv \
    && mkdir -p /app/bin/ \
    && cp ./bin/${USVC_NAME}/${USVC_NAME}_srv /app/bin/usp_${USVC_NAME}_srv

############################
# STEP 2 build a small image
############################
# using base nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM gcr.io/distroless/static
ARG USVC_NAME

# Copy our static executable
COPY --from=builder /app/bin/usp_${USVC_NAME}_srv /usr/bin/usp_${USVC_NAME}_srv

# Run the produced binary.
ENTRYPOINT ["/usr/bin/usp_${USVC_NAME}_srv"]
