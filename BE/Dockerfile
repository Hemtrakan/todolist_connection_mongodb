FROM golang:1.22-alpine AS build_stage

ENV ENV="dev"

#"DEVELOP"
#"DEVSPACE"
#"PRODUCTION"
ENV FLAVOR="DEVELOP"

RUN apk update && apk upgrade && apk add --no-cache build-base bash git openssh
LABEL maintainer="Chirapon Hemtrakan <chirapon.job@gmail.com>"

ARG PACKAGE_PATH=BE
RUN mkdir -p /go/src/
WORKDIR /go/src/${PACKAGE_PATH}

COPY . /go/src/${PACKAGE_PATH}/
RUN go mod download

RUN go build -o BE

ENTRYPOINT ./BE

EXPOSE 80

CMD ["sh", "-c", "/BE"]