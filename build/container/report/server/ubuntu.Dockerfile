############################
# STEP 1 build executable binary
############################
FROM golang:1.14.1-buster as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR $GOPATH/src/usplay/api/todo

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download \
    && go mod verify

COPY . .

# Build the binary
RUN mkdir -p /app/bin \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /app/bin/usplay_todo cmd/server/main.go

############################
# STEP 2 build an ubuntu image
############################
FROM ubuntu:20.04

# Copy our static executable
COPY --from=builder /app/bin/usplay_todo /usr/bin/usplay_todo

# Run the hello binary.
ENTRYPOINT ["/usr/bin/usplay_todo"]
