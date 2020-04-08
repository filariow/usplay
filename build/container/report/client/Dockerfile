############################
# STEP 1 build executable binary
############################
FROM golang:1.14.1-buster as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR $GOPATH/src/usplay/api/report

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download \
    && go mod verify

COPY . .

# Build the binary
RUN mkdir -p /app/bin \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /app/bin/usplay_report cmd/client/main.go

############################
# STEP 2 build a small image
############################
# using base nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM gcr.io/distroless/static

# Copy our static executable
COPY --from=builder /app/bin/usplay_report /usr/bin/usplay_report

# Run the hello binary.
ENTRYPOINT ["/usr/bin/usplay_report"]
