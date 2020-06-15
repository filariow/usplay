FROM golang:alpine
ARG USER=go

RUN apk add --no-cache --update sudo make curl gcc libc-dev

# add new user
RUN adduser -D $USER \
        && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
        && chmod 0440 /etc/sudoers.d/$USER

ARG WAIT_VERSION=2.7.3
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/${WAIT_VERSION}/wait /wait
RUN chmod +x /wait \
    && mv /wait /bin/wait

WORKDIR /app
RUN mkdir -p /app/result
COPY go.mod .
COPY go.sum .

USER $USER
RUN go mod download -x 
COPY . .
RUN sudo chown -R $USER /app

ENV RESULT_DIR /app/result
WORKDIR /app/test/integration/bookmaster
CMD sudo chown -R $(id -un) $RESULT_DIR && /bin/wait 5 && make integration-tests