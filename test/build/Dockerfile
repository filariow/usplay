ARG PROJECTS="order report activity activitytype"

############################
# STEP 1 build executable binary
############################
FROM golang:1.14.1-buster as builder
ARG PROJECTS

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
RUN chmod +x build/bin/build.sh \
        && bash build/bin/build.sh ${PROJECTS}
