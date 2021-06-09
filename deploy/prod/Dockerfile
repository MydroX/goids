FROM golang:1.16.0-alpine3.13

ARG WORK_DIR

WORKDIR $WORK_DIR

COPY . .

RUN apk add --no-cache ca-certificates \
    bash \
    sudo \
    git \
    nodejs=14.16.0-r0 \
    npm=14.16.0-r0 \
    yarn