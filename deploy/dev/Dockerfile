FROM golang:1.16.0-alpine3.13

ARG WORK_DIR

WORKDIR $WORK_DIR

COPY . .

RUN apk add --no-cache ca-certificates \
    bash \
    sudo \
    git \
    nodejs=14.16.1-r1 \
    npm=14.16.1-r1 \
    yarn=1.22.10-r0

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/mgechev/revive

ENTRYPOINT [ "/go/src/github.com/MydroX/goids/deploy/dev/entrypoint.sh" ]